# VK_NV_geometry_shader_passthrough(3) Manual Page

## Name

VK_NV_geometry_shader_passthrough - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

96

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_geometry_shader_passthrough](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_geometry_shader_passthrough.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_geometry_shader_passthrough%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_geometry_shader_passthrough%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-02-15

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_geometry_shader_passthrough`](https://registry.khronos.org/OpenGL/extensions/NV/NV_geometry_shader_passthrough.txt)

- This extension requires the `geometryShader` feature.

**Contributors**  
- Piers Daniell, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_NV_geometry_shader_passthrough`

Geometry shaders provide the ability for applications to process each
primitive sent through the graphics pipeline using a programmable
shader. However, one common use case treats them largely as a
“passthrough”. In this use case, the bulk of the geometry shader code
simply copies inputs from each vertex of the input primitive to
corresponding outputs in the vertices of the output primitive. Such
shaders might also compute values for additional built-in or
user-defined per-primitive attributes (e.g., `Layer`) to be assigned to
all the vertices of the output primitive.

This extension provides access to the `PassthroughNV` decoration under
the `GeometryShaderPassthroughNV` capability. Adding this to a geometry
shader input variable specifies that the values of this input are copied
to the corresponding vertex of the output primitive.

When using GLSL source-based shading languages, the `passthrough` layout
qualifier from `GL_NV_geometry_shader_passthrough` maps to the
`PassthroughNV` decoration. To use the `passthrough` layout, in GLSL the
`GL_NV_geometry_shader_passthrough` extension must be enabled. Behavior
is described in the `GL_NV_geometry_shader_passthrough` extension
specification.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_GEOMETRY_SHADER_PASSTHROUGH_EXTENSION_NAME`

- `VK_NV_GEOMETRY_SHADER_PASSTHROUGH_SPEC_VERSION`

## <a href="#_new_variable_decoration" class="anchor"></a>New Variable Decoration

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry-passthrough-passthrough"
  target="_blank" rel="noopener"><code>PassthroughNV</code></a> in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry-passthrough"
  target="_blank" rel="noopener">Geometry Shader Passthrough</a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-GeometryShaderPassthroughNV"
  target="_blank"
  rel="noopener"><code>GeometryShaderPassthroughNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1\) Should we require or allow a passthrough geometry shader to specify
the output layout qualifiers for the output primitive type and maximum
vertex count in the SPIR-V?

**RESOLVED**: Yes they should be required in the SPIR-V. Per
GL_NV_geometry_shader_passthrough they are not permitted in the GLSL
source shader, but SPIR-V is lower-level. It is straightforward for the
GLSL compiler to infer them from the input primitive type and to
explicitly emit them in the SPIR-V according to the following table.

| Input Layout | Implied Output Layout                    |
|--------------|------------------------------------------|
| points       | `layout(points, max_vertices=1)`         |
| lines        | `layout(line_strip, max_vertices=2)`     |
| triangles    | `layout(triangle_strip, max_vertices=3)` |

2\) How does interface matching work with passthrough geometry shaders?

**RESOLVED**: This is described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry-passthrough-interface"
target="_blank" rel="noopener">Passthrough Interface Matching</a>. In GL
when using passthrough geometry shaders in separable mode, all inputs
must also be explicitly assigned location layout qualifiers. In Vulkan
all SPIR-V shader inputs (except built-ins) must also have location
decorations specified. Redeclarations of built-in variables that add the
passthrough layout qualifier are exempted from the rule requiring
location assignment because built-in variables do not have locations and
are matched by `BuiltIn` decoration.

## <a href="#_sample_code" class="anchor"></a>Sample Code

Consider the following simple geometry shader in unextended GLSL:

``` c
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

In this shader, the inputs `gl_Position`, `Inputs.texcoord`, and
`Inputs.baseColor` are simply copied from the input vertex to the
corresponding output vertex. The only “interesting” work done by the
geometry shader is computing and emitting a `gl_Layer` value for the
primitive.

The following geometry shader, using this extension, is equivalent:

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-15 (Daniel Koch)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_geometry_shader_passthrough"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
