# VK\_KHR\_global\_priority(3) Manual Page

## Name

VK\_KHR\_global\_priority - device extension



## [](#_registered_extension_number)Registered Extension Number

189

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

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_global_priority%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_global_priority%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-10-22

**Contributors**

- Tobias Hector, AMD
- Contributors to `VK_EXT_global_priority`
- Contributors to `VK_EXT_global_priority_query`

## [](#_description)Description

In Vulkan, users can specify device-scope queue priorities. In some cases it may be useful to extend this concept to a system-wide scope. This device extension allows applications to query the global queue priorities supported by a queue family, and then set a priority when creating queues. The default queue priority is `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`.

Implementations can report which global priority levels are treated differently by the implementation. It is intended primarily for use in system integration along with certain platform-specific priority enforcement rules.

The driver implementation will attempt to skew hardware resource allocation in favor of the higher-priority task. Therefore, higher-priority work may retain similar latency and throughput characteristics even if the system is congested with lower priority work.

The global priority level of a queue shall take precedence over the per-process queue priority (`VkDeviceQueueCreateInfo`::`pQueuePriorities`).

Abuse of this feature may result in starving the rest of the system from hardware resources. Therefore, the driver implementation may deny requests to acquire a priority above the default priority (`VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`) if the caller does not have sufficient privileges. In this scenario `VK_ERROR_NOT_PERMITTED_EXT` is returned.

The driver implementation may fail the queue allocation request if resources required to complete the operation have been exhausted (either by the same process or a different process). In this scenario `VK_ERROR_INITIALIZATION_FAILED` is returned.

## [](#_new_structures)New Structures

- Extending [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html):
  
  - [VkDeviceQueueGlobalPriorityCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueGlobalPriorityCreateInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyGlobalPriorityPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyGlobalPriorityPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriorityKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_GLOBAL_PRIORITY_EXTENSION_NAME`
- `VK_KHR_GLOBAL_PRIORITY_SPEC_VERSION`
- `VK_MAX_GLOBAL_PRIORITY_SIZE_KHR`
- Extending [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html):
  
  - `VK_QUEUE_GLOBAL_PRIORITY_HIGH_KHR`
  - `VK_QUEUE_GLOBAL_PRIORITY_LOW_KHR`
  - `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR`
  - `VK_QUEUE_GLOBAL_PRIORITY_REALTIME_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_NOT_PERMITTED_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_issues)Issues

1\) Can we additionally query whether a caller is permitted to acquire a specific global queue priority in this extension?

**RESOLVED**: No. Whether a caller has enough privilege goes with the OS, and the Vulkan driver cannot really guarantee that the privilege will not change in between this query and the actual queue creation call.

2\) If more than 1 queue using global priority is requested, is there a good way to know which queue is failing the device creation?

**RESOLVED**: No. There is not a good way at this moment, and it is also not quite actionable for the applications to know that because the information may not be accurate. Queue creation can fail because of runtime constraints like insufficient privilege or lack of resource, and the failure is not necessarily tied to that particular queue configuration requested.

## [](#_version_history)Version History

- Revision 1, 2021-10-22 (Tobias Hector)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_global_priority)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0