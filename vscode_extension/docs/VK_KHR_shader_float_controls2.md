# VK\_KHR\_shader\_float\_controls2(3) Manual Page

## Name

VK\_KHR\_shader\_float\_controls2 - device extension



## [](#_registered_extension_number)Registered Extension Number

529

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_shader\_float\_controls](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float_controls.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_float\_controls2](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_float_controls2.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Graeme Leese [\[GitHub\]gnl21](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_float_controls2%5D%20%40gnl21%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_float_controls2%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_float\_controls2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_float_controls2.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-05-16

**Interactions and External Dependencies**

- This extension requires [`SPV_KHR_float_controls2`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_float_controls2.html).

**Contributors**

- Graeme Leese, Broadcom

## [](#_description)Description

This extension enables use of the more expressive fast floating-point math flags in the SPV\_KHR\_float\_controls2 extension. These flags give finer- grained control over which optimizations compilers may apply, potentially speeding up execution while retaining correct results.

The extension also adds control over the fast-math modes to the GLSL extended instruction set, making these operations more consistent with SPIR-V and allowing their use in situations where floating-point conformance is important.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderFloatControls2FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderFloatControls2FeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_FLOAT_CONTROLS_2_EXTENSION_NAME`
- `VK_KHR_SHADER_FLOAT_CONTROLS_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT_CONTROLS_2_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [FloatControls2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-FloatControls2)

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2023-05-16 (Graeme Leese)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_float_controls2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0