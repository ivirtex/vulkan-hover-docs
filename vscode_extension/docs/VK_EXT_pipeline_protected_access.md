# VK\_EXT\_pipeline\_protected\_access(3) Manual Page

## Name

VK\_EXT\_pipeline\_protected\_access - device extension



## [](#_registered_extension_number)Registered Extension Number

467

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pipeline_protected_access%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pipeline_protected_access%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_pipeline\_protected\_access](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_pipeline_protected_access.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-07-28

**Contributors**

- Shahbaz Youssefi, Google
- Jörg Wagner, Arm
- Ralph Potter, Samsung
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension allows protected memory access to be specified per pipeline as opposed to per device. Through the usage of this extension, any performance penalty paid due to access to protected memory will be limited to the specific pipelines that make such accesses.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePipelineProtectedAccessFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineProtectedAccessFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PIPELINE_PROTECTED_ACCESS_EXTENSION_NAME`
- `VK_EXT_PIPELINE_PROTECTED_ACCESS_SPEC_VERSION`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT`
  - `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_PROTECTED_ACCESS_FEATURES_EXT`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the EXT suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2022-07-28 (Shahbaz Youssefi)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_pipeline_protected_access).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0