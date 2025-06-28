# VK\_KHR\_shader\_expect\_assume(3) Manual Page

## Name

VK\_KHR\_shader\_expect\_assume - device extension



## [](#_registered_extension_number)Registered Extension Number

545

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_expect\_assume](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_expect_assume.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_expect_assume%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_expect_assume%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_expect\_assume](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_expect_assume.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-12-06

**IP Status**

No known IP claims.

**Contributors**

- Kevin Petit, Arm
- Tobias Hector, AMD
- James Fitzpatrick, Imagination Technologies

## [](#_description)Description

This extension allows the use of the `SPV_KHR_expect_assume` extension in SPIR-V shader modules which enables SPIR-V producers to provide optimization hints to the Vulkan implementation.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderExpectAssumeFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderExpectAssumeFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_EXPECT_ASSUME_EXTENSION_NAME`
- `VK_KHR_SHADER_EXPECT_ASSUME_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_EXPECT_ASSUME_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [ExpectAssumeKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ExpectAssumeKHR)

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2023-12-06 (Kevin Petit)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_expect_assume)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0