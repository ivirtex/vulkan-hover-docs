# VK_KHR_global_priority(3) Manual Page

## Name

VK_KHR_global_priority - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

189

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_global_priority%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_global_priority%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-10-22

**Contributors**  
- Tobias Hector, AMD

- Contributors to
  [`VK_EXT_global_priority`](VK_EXT_global_priority.html)

- Contributors to
  [`VK_EXT_global_priority_query`](VK_EXT_global_priority_query.html)

## <a href="#_description" class="anchor"></a>Description

In Vulkan, users can specify device-scope queue priorities. In some
cases it may be useful to extend this concept to a system-wide scope.
This device extension allows applications to query the global queue
priorities supported by a queue family, and then set a priority when
creating queues. The default queue priority is
`VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`.

Implementations can report which global priority levels are treated
differently by the implementation. It is intended primarily for use in
system integration along with certain platform-specific priority
enforcement rules.

The driver implementation will attempt to skew hardware resource
allocation in favor of the higher-priority task. Therefore,
higher-priority work may retain similar latency and throughput
characteristics even if the system is congested with lower priority
work.

The global priority level of a queue shall take precedence over the
per-process queue priority
(`VkDeviceQueueCreateInfo`::`pQueuePriorities`).

Abuse of this feature may result in starving the rest of the system from
hardware resources. Therefore, the driver implementation may deny
requests to acquire a priority above the default priority
(`VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`) if the caller does not have
sufficient privileges. In this scenario `VK_ERROR_NOT_PERMITTED_EXT` is
returned.

The driver implementation may fail the queue allocation request if
resources required to complete the operation have been exhausted (either
by the same process or a different process). In this scenario
`VK_ERROR_INITIALIZATION_FAILED` is returned.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html):

  - [VkDeviceQueueGlobalPriorityCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR.html)

- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html):

  - [VkQueueFamilyGlobalPriorityPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyGlobalPriorityPropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_GLOBAL_PRIORITY_EXTENSION_NAME`

- `VK_KHR_GLOBAL_PRIORITY_SPEC_VERSION`

- `VK_MAX_GLOBAL_PRIORITY_SIZE_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_NOT_PERMITTED_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Can we additionally query whether a caller is permitted to acquire a
specific global queue priority in this extension?

**RESOLVED**: No. Whether a caller has enough privilege goes with the
OS, and the Vulkan driver cannot really guarantee that the privilege
will not change in between this query and the actual queue creation
call.

2\) If more than 1 queue using global priority is requested, is there a
good way to know which queue is failing the device creation?

**RESOLVED**: No. There is not a good way at this moment, and it is also
not quite actionable for the applications to know that because the
information may not be accurate. Queue creation can fail because of
runtime constraints like insufficient privilege or lack of resource, and
the failure is not necessarily tied to that particular queue
configuration requested.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-10-22 (Tobias Hector)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_global_priority"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
