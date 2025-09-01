# VK\_NV\_geometry\_shader\_passthrough(3) Manual Page

## Name

VK\_NV\_geometry\_shader\_passthrough - device extension



## [](#_registered_extension_number)Registered Extension Number

96

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_geometry\_shader\_passthrough](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_geometry_shader_passthrough.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_geometry_shader_passthrough%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_geometry_shader_passthrough%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-02-15

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_geometry_shader_passthrough`](https://registry.khronos.org/OpenGL/extensions/NV/NV_geometry_shader_passthrough.txt)
- This extension requires the [`geometryShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-geometryShader) feature.

**Contributors**

- Piers Daniell, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_NV_geometry_shader_passthrough`

Geometry shaders provide the ability for applications to process each primitive sent through the graphics pipeline using a programmable shader. However, one common use case treats them largely as a “passthrough”. In this use case, the bulk of the geometry shader code simply copies inputs from each vertex of the input primitive to corresponding outputs in the vertices of the output primitive. Such shaders might also compute values for additional built-in or user-defined per-primitive attributes (e.g., `Layer`) to be assigned to all the vertices of the output primitive.

This extension provides access to the `PassthroughNV` decoration under the `GeometryShaderPassthroughNV` capability. Adding this to a geometry shader input variable specifies that the values of this input are copied to the corresponding vertex of the output primitive.

When using GLSL source-based shading languages, the `passthrough` layout qualifier from `GL_NV_geometry_shader_passthrough` maps to the `PassthroughNV` decoration. To use the `passthrough` layout, in GLSL the `GL_NV_geometry_shader_passthrough` extension must be enabled. Behavior is described in the `GL_NV_geometry_shader_passthrough` extension specification.

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_GEOMETRY_SHADER_PASSTHROUGH_EXTENSION_NAME`
- `VK_NV_GEOMETRY_SHADER_PASSTHROUGH_SPEC_VERSION`

## [](#_new_variable_decoration)New Variable Decoration

- [`PassthroughNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#geometry-passthrough-passthrough) in [Geometry Shader Passthrough](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#geometry-passthrough)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`GeometryShaderPassthroughNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-GeometryShaderPassthroughNV)

## [](#_issues)Issues

1\) Should we require or allow a passthrough geometry shader to specify the output layout qualifiers for the output primitive type and maximum vertex count in the SPIR-V?

**RESOLVED**: Yes they should be required in the SPIR-V. Per GL\_NV\_geometry\_shader\_passthrough they are not permitted in the GLSL source shader, but SPIR-V is lower-level. It is straightforward for the GLSL compiler to infer them from the input primitive type and to explicitly emit them in the SPIR-V according to the following table.

  Input Layout Implied Output Layout

points

`layout(points, max_vertices=1)`

lines

`layout(line_strip, max_vertices=2)`

triangles

`layout(triangle_strip, max_vertices=3)`

2\) How does interface matching work with passthrough geometry shaders?

**RESOLVED**: This is described in [Passthrough Interface Matching](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#geometry-passthrough-interface). In GL when using passthrough geometry shaders in separable mode, all inputs must also be explicitly assigned location layout qualifiers. In Vulkan all SPIR-V shader inputs (except built-ins) must also have location decorations specified. Redeclarations of built-in variables that add the passthrough layout qualifier are exempted from the rule requiring location assignment because built-in variables do not have locations and are matched by `BuiltIn` decoration.

## [](#_sample_code)Sample Code

Consider the following simple geometry shader in unextended GLSL:

```c
layout(triangles) in;
layout(triangle_strip) out;
layout(max_vertices=3) out;

in Inputs {
    vec2 texcoord;
    vec4 baseColor;
} v_in[];
out Outputs {
    vec2 texcoord;
    vec4 baseColor;
};

void main()
{
    int layer = compute_layer();
    for (int i = 0; i < 3; i++) {
        gl_Position = gl_in[i].gl_Position;
        texcoord = v_in[i].texcoord;
        baseColor = v_in[i].baseColor;
        gl_Layer = layer;
        EmitVertex();
    }
}
```

In this shader, the inputs `gl_Position`, `Inputs.texcoord`, and `Inputs.baseColor` are simply copied from the input vertex to the corresponding output vertex. The only “interesting” work done by the geometry shader is computing and emitting a `gl_Layer` value for the primitive.

The following geometry shader, using this extension, is equivalent:

```c
#extension GL_NV_geometry_shader_passthrough : require

layout(triangles) in;
// No output primitive layout qualifiers required.

// Redeclare gl_PerVertex to pass through "gl_Position".
layout(passthrough) in gl_PerVertex {
    vec4 gl_Position;
} gl_in[];

// Declare "Inputs" with "passthrough" to automatically copy members.
layout(passthrough) in Inputs {
    vec2 texcoord;
    vec4 baseColor;
} v_in[];

// No output block declaration required.

void main()
{
    // The shader simply computes and writes gl_Layer.  We do not
    // loop over three vertices or call EmitVertex().
    gl_Layer = compute_layer();
}
```

## [](#_version_history)Version History

- Revision 1, 2017-02-15 (Daniel Koch)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_geometry_shader_passthrough).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0