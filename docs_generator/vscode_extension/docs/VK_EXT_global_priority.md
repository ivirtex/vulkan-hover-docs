# VK\_EXT\_global\_priority(3) Manual Page

## Name

VK\_EXT\_global\_priority - device extension



## [](#_registered_extension_number)Registered Extension Number

175

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Andres Rodriguez [\[GitHub\]lostgoat](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_global_priority%5D%20%40lostgoat%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_global_priority%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-10-06

**IP Status**

No known IP claims.

**Contributors**

- Andres Rodriguez, Valve
- Pierre-Loup Griffais, Valve
- Dan Ginsburg, Valve
- Mitch Singer, AMD

## [](#_description)Description

In Vulkan, users can specify device-scope queue priorities. In some cases it may be useful to extend this concept to a system-wide scope. This extension provides a mechanism for callers to set their system-wide priority. The default queue priority is `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`.

The driver implementation will attempt to skew hardware resource allocation in favor of the higher-priority task. Therefore, higher-priority work may retain similar latency and throughput characteristics even if the system is congested with lower priority work.

The global priority level of a queue shall take precedence over the per-process queue priority (`VkDeviceQueueCreateInfo`::`pQueuePriorities`).

Abuse of this feature may result in starving the rest of the system from hardware resources. Therefore, the driver implementation may deny requests to acquire a priority above the default priority (`VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`) if the caller does not have sufficient privileges. In this scenario `VK_ERROR_NOT_PERMITTED_EXT` is returned.

The driver implementation may fail the queue allocation request if resources required to complete the operation have been exhausted (either by the same process or a different process). In this scenario `VK_ERROR_INITIALIZATION_FAILED` is returned.

## [](#_new_structures)New Structures

- Extending [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html):
  
  - [VkDeviceQueueGlobalPriorityCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueGlobalPriorityCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkQueueGlobalPriorityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriorityEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_GLOBAL_PRIORITY_EXTENSION_NAME`
- `VK_EXT_GLOBAL_PRIORITY_SPEC_VERSION`
- Extending [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html):
  
  - `VK_QUEUE_GLOBAL_PRIORITY_HIGH_EXT`
  - `VK_QUEUE_GLOBAL_PRIORITY_LOW_EXT`
  - `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`
  - `VK_QUEUE_GLOBAL_PRIORITY_REALTIME_EXT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_NOT_PERMITTED_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 2, 2017-11-03 (Andres Rodriguez)
  
  - Fixed VkQueueGlobalPriorityEXT missing \_EXT suffix
- Revision 1, 2017-10-06 (Andres Rodriguez)
  
  - First version.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_global_priority)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0