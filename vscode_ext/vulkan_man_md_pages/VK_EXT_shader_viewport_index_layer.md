# VK_EXT_shader_viewport_index_layer(3) Manual Page

## Name

VK_EXT_shader_viewport_index_layer - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

163

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_shader_viewport_index_layer](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_shader_viewport_index_layer.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_viewport_index_layer%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_viewport_index_layer%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-08-08

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_shader_viewport_layer_array`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_viewport_layer_array.txt),
  [`GL_AMD_vertex_shader_layer`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_vertex_shader_layer.txt),
  [`GL_AMD_vertex_shader_viewport_index`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_vertex_shader_viewport_index.txt),
  and
  [`GL_NV_viewport_array2`](https://registry.khronos.org/OpenGL/extensions/NV/NV_viewport_array2.txt)

- This extension requires the `multiViewport` feature.

- This extension interacts with the `tessellationShader` feature.

**Contributors**  
- Piers Daniell, NVIDIA

- Jeff Bolz, NVIDIA

- Jan-Harald Fredriksen, ARM

- Daniel Rakos, AMD

- Slawomir Grajeswki, Intel

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the `ShaderViewportIndexLayerEXT`
capability from the `SPV_EXT_shader_viewport_index_layer` extension in
Vulkan.

This extension allows variables decorated with the `Layer` and
`ViewportIndex` built-ins to be exported from vertex or tessellation
shaders, using the `ShaderViewportIndexLayerEXT` capability.

When using GLSL source-based shading languages, the `gl_ViewportIndex`
and `gl_Layer` built-in variables map to the SPIR-V `ViewportIndex` and
`Layer` built-in decorations, respectively. Behavior of these variables
is extended as described in the `GL_ARB_shader_viewport_layer_array` (or
the precursor `GL_AMD_vertex_shader_layer`,
`GL_AMD_vertex_shader_viewport_index`, and `GL_NV_viewport_array2`
extensions).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>ShaderViewportIndexLayerEXT</code> capability is equivalent
to the <code>ShaderViewportIndexLayerNV</code> capability added by <a
href="VK_NV_viewport_array2.html"><code>VK_NV_viewport_array2</code></a>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2.

The single `ShaderViewportIndexLayerEXT` capability from the
`SPV_EXT_shader_viewport_index_layer` extension is replaced by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndex"
target="_blank" rel="noopener"><code>ShaderViewportIndex</code></a> and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderLayer"
target="_blank" rel="noopener"><code>ShaderLayer</code></a> capabilities
from SPIR-V 1.5 which are enabled by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderOutputViewportIndex"
target="_blank"
rel="noopener"><code>shaderOutputViewportIndex</code></a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderOutputLayer"
target="_blank" rel="noopener"><code>shaderOutputLayer</code></a>
features, respectively. Additionally, if Vulkan 1.2 is supported but
this extension is not, these capabilities are optional.

Enabling both features is equivalent to enabling the
`VK_EXT_shader_viewport_index_layer` extension.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_EXTENSION_NAME`

- `VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_SPEC_VERSION`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- (modified) <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-layer"
  target="_blank" rel="noopener"><code>Layer</code></a>

- (modified) <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-viewportindex"
  target="_blank" rel="noopener"><code>ViewportIndex</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndexLayerEXT"
  target="_blank"
  rel="noopener"><code>ShaderViewportIndexLayerEXT</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-08-08 (Daniel Koch)

  - Internal drafts

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_viewport_index_layer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
