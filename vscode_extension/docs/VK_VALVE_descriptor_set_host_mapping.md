# VK\_VALVE\_descriptor\_set\_host\_mapping(3) Manual Page

## Name

VK\_VALVE\_descriptor\_set\_host\_mapping - device extension



## [](#_registered_extension_number)Registered Extension Number

421

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Hans-Kristian Arntzen [\[GitHub\]HansKristian-Work](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_VALVE_descriptor_set_host_mapping%5D%20%40HansKristian-Work%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_VALVE_descriptor_set_host_mapping%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-02-22

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, Valve

## [](#_description)Description

This extension allows applications to directly query a host pointer for a [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) which **can** be used to copy descriptors between descriptor sets without the use of an API command. Memory offsets and sizes for descriptors **can** be queried from a [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) as well.

Note

There is currently no specification language written for this extension. The links to APIs defined by the extension are to stubs that only include generated content such as API declarations and implicit valid usage statements.

Note

This extension is only intended for use in specific embedded environments with known implementation details, and is therefore undocumented.

## [](#_new_commands)New Commands

- [vkGetDescriptorSetHostMappingVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetHostMappingVALVE.html)
- [vkGetDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutHostMappingInfoVALVE.html)

## [](#_new_structures)New Structures

- [VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetBindingReferenceVALVE.html)
- [VkDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutHostMappingInfoVALVE.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_VALVE_DESCRIPTOR_SET_HOST_MAPPING_EXTENSION_NAME`
- `VK_VALVE_DESCRIPTOR_SET_HOST_MAPPING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_BINDING_REFERENCE_VALVE`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_HOST_MAPPING_INFO_VALVE`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_SET_HOST_MAPPING_FEATURES_VALVE`

## [](#_version_history)Version History

- Revision 1, 2022-02-22 (Hans-Kristian Arntzen)
  
  - Initial specification

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_VALVE_descriptor_set_host_mapping)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0