# VK\_KHR\_maintenance4(3) Manual Page

## Name

VK\_KHR\_maintenance4 - device extension



## [](#_registered_extension_number)Registered Extension Number

414

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance4%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance4%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-10-25

**Interactions and External Dependencies**

- Requires SPIR-V 1.2 for `LocalSizeId`

**Contributors**

- Lionel Duc, NVIDIA
- Faith Ekstrand, Intel
- Spencer Fricke, Samsung
- Tobias Hector, AMD
- Lionel Landwerlin, Intel
- Graeme Leese, Broadcom
- Tom Olson, Arm
- Stu Smith, AMD
- Yiwei Zhang, Google

## [](#_description)Description

`VK_KHR_maintenance4` adds a collection of minor features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Allow the application to destroy their [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object immediately after it was used to create another object. It is no longer necessary to keep its handle valid while the created object is in use.
- Add a new [`maxBufferSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxBufferSize) implementation-defined limit for the maximum size `VkBuffer` that **can** be created.
- Add support for the SPIR-V 1.2 `LocalSizeId` execution mode, which can be used as an alternative to `LocalSize` to specify the local workgroup size with specialization constants.
- Add a guarantee that images created with identical creation parameters will always have the same alignment requirements.
- Add new [vkGetDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html), [vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirementsKHR.html), and [vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html) to allow the application to query the image memory requirements without having to create an image object and query it.
- Relax the requirement that push constants must be initialized before they are dynamically accessed.
- Relax the interface matching rules to allow a larger output vector to match with a smaller input vector, with additional values being discarded.
- Add a guarantee for buffer memory requirement that the size memory requirement is never greater than the result of aligning create size with the alignment memory requirement.

## [](#_new_commands)New Commands

- [vkGetDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html)
- [vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirementsKHR.html)
- [vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html)

## [](#_new_structures)New Structures

- [VkDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceBufferMemoryRequirementsKHR.html)
- [VkDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirementsKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMaintenance4FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance4FeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMaintenance4PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance4PropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE_4_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_4_SPEC_VERSION`
- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html):
  
  - `VK_IMAGE_ASPECT_NONE_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_PROPERTIES_KHR`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2021-08-18 (Piers Daniell)
  
  - Internal revisions
- Revision 2, 2021-10-25 (Yiwei Zhang)
  
  - More guarantees on buffer memory requirements

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance4)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0