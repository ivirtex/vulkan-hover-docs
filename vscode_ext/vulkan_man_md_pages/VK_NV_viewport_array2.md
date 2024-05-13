# VK_NV_viewport_array2(3) Manual Page

## Name

VK_NV_viewport_array2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

97

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_viewport_array2](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_viewport_array2.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_viewport_array2%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_viewport_array2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-02-15

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_viewport_array2`](https://registry.khronos.org/OpenGL/extensions/NV/NV_viewport_array2.txt)

- This extension requires the `geometryShader` and `multiViewport`
  features.

- This extension interacts with the `tessellationShader` feature.

**Contributors**  
- Piers Daniell, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_NV_viewport_array2`

which allows a single primitive to be broadcast to multiple viewports
and/or multiple layers. A new shader built-in output `ViewportMaskNV` is
provided, which allows a single primitive to be output to multiple
viewports simultaneously. Also, a new SPIR-V decoration is added to
control whether the effective viewport index is added into the variable
decorated with the `Layer` built-in decoration. These capabilities allow
a single primitive to be output to multiple layers simultaneously.

This extension allows variables decorated with the `Layer` and
`ViewportIndex` built-ins to be exported from vertex or tessellation
shaders, using the `ShaderViewportIndexLayerNV` capability.

This extension adds a new `ViewportMaskNV` built-in decoration that is
available for output variables in vertex, tessellation evaluation, and
geometry shaders, and a new `ViewportRelativeNV` decoration that can be
added on variables decorated with `Layer` when using the
`ShaderViewportMaskNV` capability.

When using GLSL source-based shading languages, the
`gl_ViewportMask`\[\] built-in output variable and `viewport_relative`
layout qualifier from `GL_NV_viewport_array2` map to the
`ViewportMaskNV` and `ViewportRelativeNV` decorations, respectively.
Behavior is described in the `GL_NV_viewport_array2` extension
specification.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>ShaderViewportIndexLayerNV</code> capability is equivalent
to the <code>ShaderViewportIndexLayerEXT</code> capability added by <a
href="VK_EXT_shader_viewport_index_layer.html"><code>VK_EXT_shader_viewport_index_layer</code></a>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_VIEWPORT_ARRAY2_EXTENSION_NAME`

- `VK_NV_VIEWPORT_ARRAY2_SPEC_VERSION`

- `VK_NV_VIEWPORT_ARRAY_2_EXTENSION_NAME`

- `VK_NV_VIEWPORT_ARRAY_2_SPEC_VERSION`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- (modified) <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-layer"
  target="_blank" rel="noopener"><code>Layer</code></a>

- (modified) <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-viewportindex"
  target="_blank" rel="noopener"><code>ViewportIndex</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-viewportmask"
  target="_blank" rel="noopener"><code>ViewportMaskNV</code></a>

## <a href="#_new_variable_decoration" class="anchor"></a>New Variable Decoration

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-layer"
  target="_blank" rel="noopener"><code>ViewportRelativeNV</code> in
  <code>Layer</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndexLayerNV"
  target="_blank"
  rel="noopener"><code>ShaderViewportIndexLayerNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportMaskNV"
  target="_blank" rel="noopener"><code>ShaderViewportMaskNV</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-15 (Daniel Koch)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_viewport_array2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
