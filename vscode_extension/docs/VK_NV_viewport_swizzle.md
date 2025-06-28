# VK\_NV\_viewport\_swizzle(3) Manual Page

## Name

VK\_NV\_viewport\_swizzle - device extension



## [](#_registered_extension_number)Registered Extension Number

99

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_viewport_swizzle%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_viewport_swizzle%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-12-22

**Interactions and External Dependencies**

- This extension requires `multiViewport` and `geometryShader` features to be useful.

**Contributors**

- Daniel Koch, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension provides a new per-viewport swizzle that can modify the position of primitives sent to each viewport. New viewport swizzle state is added for each viewport, and a new position vector is computed for each vertex by selecting from and optionally negating any of the four components of the original position vector.

This new viewport swizzle is useful for a number of algorithms, including single-pass cube map rendering (broadcasting a primitive to multiple faces and reorienting the vertex position for each face) and voxel rasterization. The per-viewport component remapping and negation provided by the swizzle allows application code to re-orient three-dimensional geometry with a view along any of the **X**, **Y**, or **Z** axes. If a perspective projection and depth buffering is required, 1/W buffering should be used, as described in the single-pass cube map rendering example in the “Issues” section below.

## [](#_new_structures)New Structures

- [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html)
- Extending [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html):
  
  - [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)

## [](#_new_enums)New Enums

- [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportCoordinateSwizzleNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineViewportSwizzleStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_VIEWPORT_SWIZZLE_EXTENSION_NAME`
- `VK_NV_VIEWPORT_SWIZZLE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV`

## [](#_issues)Issues

1\) Where does viewport swizzling occur in the pipeline?

**RESOLVED**: Despite being associated with the viewport, viewport swizzling must happen prior to the viewport transform. In particular, it needs to be performed before clipping and perspective division.

The viewport mask expansion (`VK_NV_viewport_array2`) and the viewport swizzle could potentially be performed before or after transform feedback, but feeding back several viewports worth of primitives with different swizzles does not seem particularly useful. This specification applies the viewport mask and swizzle after transform feedback, and makes primitive queries only count each primitive once.

2\) Any interesting examples of how this extension, `VK_NV_viewport_array2`, and `VK_NV_geometry_shader_passthrough` can be used together in practice?

**RESOLVED**: One interesting use case for this extension is for single-pass rendering to a cube map. In this example, the application would attach a cube map texture to a layered FBO where the six cube faces are treated as layers. Vertices are sent through the vertex shader without applying a projection matrix, where the `gl_Position` output is (x,y,z,1) and the center of the cube map is at (0,0,0). With unextended Vulkan, one could have a conventional instanced geometry shader that looks something like the following:

```c
layout(invocations = 6) in;     // separate invocation per face
layout(triangles) in;
layout(triangle_strip) out;
layout(max_vertices = 3) out;

in Inputs {
vec2 texcoord;
vec3 normal;
vec4 baseColor;
} v[];

    out Outputs {
    vec2 texcoord;
    vec3 normal;
    vec4 baseColor;
    };

    void main()
    {
    int face = gl_InvocationID;  // which face am I?

    // Project gl_Position for each vertex onto the cube map face.
    vec4 positions[3];
    for (int i = 0; i < 3; i++) {
        positions[i] = rotate(gl_in[i].gl_Position, face);
    }

    // If the primitive does not project onto this face, we are done.
    if (shouldCull(positions)) {
        return;
    }

    // Otherwise, emit a copy of the input primitive to the
    // appropriate face (using gl_Layer).
    for (int i = 0; i < 3; i++) {
        gl_Layer = face;
        gl_Position = positions[i];
        texcoord = v[i].texcoord;
        normal = v[i].normal;
        baseColor = v[i].baseColor;
        EmitVertex();
    }
}
```

With passthrough geometry shaders, this can be done using a much simpler shader:

```c
layout(triangles) in;
layout(passthrough) in Inputs {
    vec2 texcoord;
    vec3 normal;
    vec4 baseColor;
}
layout(passthrough) in gl_PerVertex {
    vec4 gl_Position;
} gl_in[];
layout(viewport_relative) out int gl_Layer;

void main()
{
    // Figure out which faces the primitive projects onto and
    // generate a corresponding viewport mask.
    uint mask = 0;
    for (int i = 0; i < 6; i++) {
        if (!shouldCull(face)) {
        mask |= 1U << i;
        }
    }
    gl_ViewportMask = mask;
    gl_Layer = 0;
}
```

The application code is set up so that each of the six cube faces has a separate viewport (numbered 0 to 5). Each face also has a separate swizzle, programmed via the [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html) pipeline state. The viewport swizzle feature performs the coordinate transformation handled by the `rotate`() function in the original shader. The `viewport_relative` layout qualifier says that the viewport number (0 to 5) is added to the base `gl_Layer` value of 0 to determine which layer (cube face) the primitive should be sent to.

Note that the use of the passed through input `normal` in this example suggests that the fragment shader in this example would perform an operation like per-fragment lighting. The viewport swizzle would transform the position to be face-relative, but `normal` would remain in the original coordinate system. It seems likely that the fragment shader in either version of the example would want to perform lighting in the original coordinate system. It would likely do this by reconstructing the position of the fragment in the original coordinate system using `gl_FragCoord`, a constant or uniform holding the size of the cube face, and the input `gl_ViewportIndex` (or `gl_Layer`), which identifies the cube face. Since the value of `normal` is in the original coordinate system, it would not need to be modified as part of this coordinate transformation.

Note that while the `rotate`() operation in the regular geometry shader above could include an arbitrary post-rotation projection matrix, the viewport swizzle does not support arbitrary math. To get proper projection, 1/W buffering should be used. To do this:

1. Program the viewport swizzles to move the pre-projection W eye coordinate (typically 1.0) into the Z coordinate of the swizzle output and the eye coordinate component used for depth into the W coordinate. For example, the viewport corresponding to the +Z face might use a swizzle of (+X, -Y, +W, +Z). The Z normalized device coordinate computed after swizzling would then be z'/w' = 1/Zeye.
2. On NVIDIA implementations supporting floating-point depth buffers with values outside \[0,1], prevent unwanted near plane clipping by enabling `depthClampEnable`. Ensure that the depth clamp does not mess up depth testing by programming the depth range to very large values, such as `minDepthBounds`=-z, `maxDepthBounds`=+z, where z = 2127. It should be possible to use IEEE infinity encodings also (`0xFF800000` for `-INF`, `0x7F800000` for `+INF`). Even when near/far clipping is disabled, primitives extending behind the eye will still be clipped because one or more vertices will have a negative W coordinate and fail X/Y clipping tests.
   
   On other implementations, scale X, Y, and Z eye coordinates so that vertices on the near plane have a post-swizzle W coordinate of 1.0. For example, if the near plane is at Zeye = 1/256, scale X, Y, and Z by 256.
3. Adjust depth testing to reflect the fact that 1/W values are large near the eye and small away from the eye. Clear the depth buffer to zero (infinitely far away) and use a depth test of `VK_COMPARE_OP_GREATER` instead of `VK_COMPARE_OP_LESS`.

## [](#_version_history)Version History

- Revision 1, 2016-12-22 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_viewport_swizzle)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0