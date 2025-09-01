# VK\_NV\_per\_stage\_descriptor\_set(3) Manual Page

## Name

VK\_NV\_per\_stage\_descriptor\_set - device extension



## [](#_registered_extension_number)Registered Extension Number

517

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html)  
or  
[Vulkan Version 1.4](#versions-1.4)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_per_stage_descriptor_set%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_per_stage_descriptor_set%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-10-16

**IP Status**

No known IP claims.

**Contributors**

- Daniel Story, Nintendo

## [](#_description)Description

This extension introduces a new descriptor set layout creation flag that allows bindings in a descriptor set to be scoped to each shader stage. This means that shaders bound at the same time **may** use completely different descriptor set layouts without any restrictions on compatibility, and that the descriptor limits that would otherwise apply to the union of all stages together instead apply to each stage individually. It also means that descriptors shared by multiple stages **must** be bound to each stage or set of stages that use a unique descriptor set layout using their specific per stage descriptor set layout(s).

This extension also allows each of the new descriptor binding functions from VK\_KHR\_maintenance6 to have their [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) member be optionally set to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), in which case the pipeline layout information is taken from a [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure in the `pNext` chain. This enables descriptors to be directly bound using descriptor set layouts without applications needing to create and manage [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) objects at command recording time.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePerStageDescriptorSetFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePerStageDescriptorSetFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_PER_STAGE_DESCRIPTOR_SET_EXTENSION_NAME`
- `VK_NV_PER_STAGE_DESCRIPTOR_SET_SPEC_VERSION`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PER_STAGE_DESCRIPTOR_SET_FEATURES_NV`

## [](#_issues)Issues

None

## [](#_version_history)Version History

- Revision 1, 2023-10-16 (Piers Daniell)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_per_stage_descriptor_set).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0