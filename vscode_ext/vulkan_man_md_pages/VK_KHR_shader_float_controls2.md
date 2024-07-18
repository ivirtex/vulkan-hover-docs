# VK_KHR_shader_float_controls2(3) Manual Page

## Name

VK_KHR_shader_float_controls2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

529

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  
and  
[VK_KHR_shader_float_controls](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float_controls.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_float_controls2](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_float_controls2.html)

## <a href="#_contact" class="anchor"></a>Contact

- Graeme Leese <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_float_controls2%5D%20@gnl21%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_float_controls2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gnl21</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_shader_float_controls2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_float_controls2.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-05-16

**Interactions and External Dependencies**  
- This extension requires
  [`SPV_KHR_float_controls2`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_float_controls2.html).

**Contributors**  
- Graeme Leese, Broadcom

## <a href="#_description" class="anchor"></a>Description

This extension enables use of the more expressive fast floating-point
math flags in the SPV_KHR_float_controls2 extension. These flags give
finer- grained control over which optimizations compilers may apply,
potentially speeding up execution while retaining correct results.

The extension also adds control over the fast-math modes to the GLSL
extended instruction set, making these operations more consistent with
SPIR-V and allowing their use in situations where floating-point
conformance is important.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderFloatControls2FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloatControls2FeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_FLOAT_CONTROLS_2_EXTENSION_NAME`

- `VK_KHR_SHADER_FLOAT_CONTROLS_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT_CONTROLS_2_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FloatControls2"
  target="_blank" rel="noopener">FloatControls2</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-05-16 (Graeme Leese)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_float_controls2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
