# VK\_NV\_viewport\_array2(3) Manual Page

## Name

VK\_NV\_viewport\_array2 - device extension



## [](#_registered_extension_number)Registered Extension Number

97

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_viewport\_array2](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_viewport_array2.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_viewport_array2%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_viewport_array2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-02-15

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_viewport_array2`](https://registry.khronos.org/OpenGL/extensions/NV/NV_viewport_array2.txt)
- This extension requires the [`geometryShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-geometryShader) and [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) features.
- This extension interacts with the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature.

**Contributors**

- Piers Daniell, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_NV_viewport_array2`

which allows a single primitive to be broadcast to multiple viewports and/or multiple layers. A new shader built-in output `ViewportMaskNV` is provided, which allows a single primitive to be output to multiple viewports simultaneously. Also, a new SPIR-V decoration is added to control whether the effective viewport index is added into the variable decorated with the `Layer` built-in decoration. These capabilities allow a single primitive to be output to multiple layers simultaneously.

This extension allows variables decorated with the `Layer` and `ViewportIndex` built-ins to be exported from vertex or tessellation shaders, using the `ShaderViewportIndexLayerNV` capability.

This extension adds a new `ViewportMaskNV` built-in decoration that is available for output variables in vertex, tessellation evaluation, and geometry shaders, and a new `ViewportRelativeNV` decoration that can be added on variables decorated with `Layer` when using the `ShaderViewportMaskNV` capability.

When using GLSL source-based shading languages, the `gl_ViewportMask`\[] built-in output variable and `viewport_relative` layout qualifier from `GL_NV_viewport_array2` map to the `ViewportMaskNV` and `ViewportRelativeNV` decorations, respectively. Behavior is described in the `GL_NV_viewport_array2` extension specification.

Note

The `ShaderViewportIndexLayerNV` capability is equivalent to the `ShaderViewportIndexLayerEXT` capability added by `VK_EXT_shader_viewport_index_layer`.

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_VIEWPORT_ARRAY2_EXTENSION_NAME`
- `VK_NV_VIEWPORT_ARRAY2_SPEC_VERSION`
- `VK_NV_VIEWPORT_ARRAY_2_EXTENSION_NAME`
- `VK_NV_VIEWPORT_ARRAY_2_SPEC_VERSION`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- (modified) [`Layer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-layer)
- (modified) [`ViewportIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-viewportindex)
- [`ViewportMaskNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-viewportmask)

## [](#_new_variable_decoration)New Variable Decoration

- [`ViewportRelativeNV` in `Layer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-layer)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`ShaderViewportIndexLayerNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndexLayerEXT)
- [`ShaderViewportMaskNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportMaskNV)

## [](#_version_history)Version History

- Revision 1, 2017-02-15 (Daniel Koch)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_viewport_array2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0