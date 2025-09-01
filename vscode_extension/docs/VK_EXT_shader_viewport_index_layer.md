# VK\_EXT\_shader\_viewport\_index\_layer(3) Manual Page

## Name

VK\_EXT\_shader\_viewport\_index\_layer - device extension



## [](#_registered_extension_number)Registered Extension Number

163

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_shader\_viewport\_index\_layer](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_viewport_index_layer.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_viewport_index_layer%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_viewport_index_layer%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-08-08

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_shader_viewport_layer_array`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_viewport_layer_array.txt), [`GL_AMD_vertex_shader_layer`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_vertex_shader_layer.txt), [`GL_AMD_vertex_shader_viewport_index`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_vertex_shader_viewport_index.txt), and [`GL_NV_viewport_array2`](https://registry.khronos.org/OpenGL/extensions/NV/NV_viewport_array2.txt)
- This extension requires the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature.
- This extension interacts with the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature.

**Contributors**

- Piers Daniell, NVIDIA
- Jeff Bolz, NVIDIA
- Jan-Harald Fredriksen, ARM
- Daniel Rakos, AMD
- Slawomir Grajeswki, Intel

## [](#_description)Description

This extension adds support for the `ShaderViewportIndexLayerEXT` capability from the `SPV_EXT_shader_viewport_index_layer` extension in Vulkan.

This extension allows variables decorated with the `Layer` and `ViewportIndex` built-ins to be exported from vertex or tessellation shaders, using the `ShaderViewportIndexLayerEXT` capability.

When using GLSL source-based shading languages, the `gl_ViewportIndex` and `gl_Layer` built-in variables map to the SPIR-V `ViewportIndex` and `Layer` built-in decorations, respectively. Behavior of these variables is extended as described in the `GL_ARB_shader_viewport_layer_array` (or the precursor `GL_AMD_vertex_shader_layer`, `GL_AMD_vertex_shader_viewport_index`, and `GL_NV_viewport_array2` extensions).

Note

The `ShaderViewportIndexLayerEXT` capability is equivalent to the `ShaderViewportIndexLayerNV` capability added by `VK_NV_viewport_array2`.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2.

The single `ShaderViewportIndexLayerEXT` capability from the `SPV_EXT_shader_viewport_index_layer` extension is replaced by the [`ShaderViewportIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndex) and [`ShaderLayer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderLayer) capabilities from SPIR-V 1.5 which are enabled by the [`shaderOutputViewportIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderOutputViewportIndex) and [`shaderOutputLayer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderOutputLayer) features, respectively. Additionally, if Vulkan 1.2 is supported but this extension is not, these capabilities are optional.

Enabling both features is equivalent to enabling the `VK_EXT_shader_viewport_index_layer` extension.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_EXTENSION_NAME`
- `VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_SPEC_VERSION`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- (modified) [`Layer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-layer)
- (modified) [`ViewportIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-viewportindex)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`ShaderViewportIndexLayerEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndexLayerEXT)

## [](#_version_history)Version History

- Revision 1, 2017-08-08 (Daniel Koch)
  
  - Internal drafts

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_viewport_index_layer).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0