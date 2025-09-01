# VK\_KHR\_push\_descriptor(3) Manual Page

## Name

VK\_KHR\_push\_descriptor - device extension



## [](#_registered_extension_number)Registered Extension Number

81

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_1
- Interacts with VK\_KHR\_descriptor\_update\_template

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_push_descriptor%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_push_descriptor%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-12

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Michael Worcester, Imagination Technologies

## [](#_description)Description

This extension allows descriptors to be written into the command buffer, while the implementation is responsible for managing their memory. Push descriptors may enable easier porting from older APIs and in some cases can be more efficient than writing descriptors into descriptor sets.

## [](#_new_commands)New Commands

- [vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetKHR.html)

If [Vulkan Version 1.1](#versions-1.1) or [VK\_KHR\_descriptor\_update\_template](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_descriptor_update_template.html) is supported:

- [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDevicePushDescriptorPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePushDescriptorPropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PUSH_DESCRIPTOR_EXTENSION_NAME`
- `VK_KHR_PUSH_DESCRIPTOR_SPEC_VERSION`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR`

If [Vulkan Version 1.1](#versions-1.1) or [VK\_KHR\_descriptor\_update\_template](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_descriptor_update_template.html) is supported:

- Extending [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateType.html):
  
  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2016-10-15 (Jeff Bolz)
  
  - Internal revisions
- Revision 2, 2017-09-12 (Tobias Hector)
  
  - Added interactions with Vulkan 1.1

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_push_descriptor).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0