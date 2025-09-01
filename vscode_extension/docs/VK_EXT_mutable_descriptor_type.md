# VK\_EXT\_mutable\_descriptor\_type(3) Manual Page

## Name

VK\_EXT\_mutable\_descriptor\_type - device extension



## [](#_registered_extension_number)Registered Extension Number

495

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_maintenance3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance3.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Joshua Ashton [\[GitHub\]Joshua-Ashton](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_mutable_descriptor_type%5D%20%40Joshua-Ashton%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_mutable_descriptor_type%20extension%2A)
- Hans-Kristian Arntzen [\[GitHub\]HansKristian-Work](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_mutable_descriptor_type%5D%20%40HansKristian-Work%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_mutable_descriptor_type%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_mutable\_descriptor\_type](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_mutable_descriptor_type.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-08-22

**IP Status**

No known IP claims.

**Contributors**

- Joshua Ashton, Valve
- Hans-Kristian Arntzen, Valve

## [](#_description)Description

This extension allows applications to reduce descriptor memory footprint by allowing a descriptor to be able to mutate to a given list of descriptor types depending on which descriptor types are written into, or copied into a descriptor set.

The main use case this extension intends to address is descriptor indexing with `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` where the descriptor types are completely generic, as this means applications can allocate one large descriptor set, rather than having one large descriptor set per descriptor type, which significantly bloats descriptor memory usage and causes performance issues.

This extension also adds a mechanism to declare that a descriptor pool, and therefore the descriptor sets that are allocated from it, reside only in host memory; as such these descriptors can only be updated/copied, but not bound.

These features together allow much more efficient emulation of the raw D3D12 binding model. This extension is primarily intended to be useful for API layering efforts.

## [](#_new_structures)New Structures

- [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html)
- Extending [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html), [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html):
  
  - [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_MUTABLE_DESCRIPTOR_TYPE_EXTENSION_NAME`
- `VK_EXT_MUTABLE_DESCRIPTOR_TYPE_SPEC_VERSION`
- Extending [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MUTABLE_DESCRIPTOR_TYPE_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MUTABLE_DESCRIPTOR_TYPE_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-08-22 (Jon Leech)
  
  - Initial version, promoted from [VK\_VALVE\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_mutable_descriptor_type.html).

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_mutable_descriptor_type).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0