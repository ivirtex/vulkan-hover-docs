# VK\_NV\_external\_compute\_queue(3) Manual Page

## Name

VK\_NV\_external\_compute\_queue - device extension



## [](#_registered_extension_number)Registered Extension Number

557

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Chris Lentini [\[GitHub\]clentini](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_compute_queue%5D%20%40clentini%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_compute_queue%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_external\_compute\_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_external_compute_queue.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-03-24

**Contributors**

- Chris Lentini, NVIDIA
- Eric Werness, NVIDIA
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Liam Middlebrook, NVIDIA
- Lionel Duc, NVIDIA

## [](#_description)Description

This extension gives applications the ability to join compatible external compute APIs to a `VkDevice`. In this way, the extension allows an application to achieve simultaneous execution between work submitted from these compatible external APIs and work that has been submitted through the Vulkan API.

At device creation time, an application **must** supply a [VkExternalComputeQueueDeviceCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueDeviceCreateInfoNV.html). This communicates to the implementation the maximum number of external queues that the application **can** create at once. This information **may** be used by the implementation to aid in decisions made during device creation.

After device creation, the function [vkCreateExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExternalComputeQueueNV.html) is used by an application to create a new `VkExternalComputeQueueNV` object. The `VkExternalComputeQueueNV` object holds information and reserves resources necessary for a compatible external API to be able to join a `VkDevice`. This information can be queried through the [vkGetExternalComputeQueueDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExternalComputeQueueDataNV.html) function, returning an opaque blob of data which can be passed to compatible external APIs. The application **must** finally call [vkDestroyExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyExternalComputeQueueNV.html) when it is done in order to release the reserved resources.

This extension introduces a new properties structure, [VkPhysicalDeviceExternalComputeQueuePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalComputeQueuePropertiesNV.html), which can be queried through [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html). The structure provides information on functional limits to the extension as well as a way of querying the size of the application allocated memory which **must** be passed to the [vkGetExternalComputeQueueDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExternalComputeQueueDataNV.html) function.

When creating a `VkExternalComputeQueueNV` through [vkCreateExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExternalComputeQueueNV.html), the [VkExternalComputeQueueCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueCreateInfoNV.html) structure requires an application to supply a `VkQueue` to aid in external compute queue creation. The supplied `VkQueue` is a strong scheduling hint about which queue it expects to submit graphics workloads to and with which it expects simultaneous execution of compute workloads submitted through the external API.

## [](#_new_object_types)New Object Types

- [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html)

## [](#_new_commands)New Commands

- [vkCreateExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExternalComputeQueueNV.html)
- [vkDestroyExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyExternalComputeQueueNV.html)
- [vkGetExternalComputeQueueDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExternalComputeQueueDataNV.html)

## [](#_new_structures)New Structures

- [VkExternalComputeQueueCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueCreateInfoNV.html)
- [VkExternalComputeQueueDataParamsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueDataParamsNV.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkExternalComputeQueueDeviceCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueDeviceCreateInfoNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceExternalComputeQueuePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalComputeQueuePropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_EXTERNAL_COMPUTE_QUEUE_EXTENSION_NAME`
- `VK_NV_EXTERNAL_COMPUTE_QUEUE_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_EXTERNAL_COMPUTE_QUEUE_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXTERNAL_COMPUTE_QUEUE_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_EXTERNAL_COMPUTE_QUEUE_DATA_PARAMS_NV`
  - `VK_STRUCTURE_TYPE_EXTERNAL_COMPUTE_QUEUE_DEVICE_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_COMPUTE_QUEUE_PROPERTIES_NV`

Note

While the external queue is now a part of a `VkDevice`, idling the device through [vkDeviceWaitIdle](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDeviceWaitIdle.html) does not wait for the external queue. Draining the work on an external queue **must** be done through its own external API. External queues **must** be idled before destroying the associated `VkDevice`.

Note

In general, synchronization and resource sharing between the external API and Vulkan must still be accomplished via existing cross-API interop mechanisms.

## [](#_issues)Issues

## [](#_version_history)Version History

- Revision 1, 2024-05-20 (Chris Lentini)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_external_compute_queue)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0