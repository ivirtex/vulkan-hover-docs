# VK_EXT_global_priority(3) Manual Page

## Name

VK_EXT_global_priority - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

175

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to [VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Andres Rodriguez <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_global_priority%5D%20@lostgoat%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_global_priority%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>lostgoat</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-10-06

**IP Status**  
No known IP claims.

**Contributors**  
- Andres Rodriguez, Valve

- Pierre-Loup Griffais, Valve

- Dan Ginsburg, Valve

- Mitch Singer, AMD

## <a href="#_description" class="anchor"></a>Description

In Vulkan, users can specify device-scope queue priorities. In some
cases it may be useful to extend this concept to a system-wide scope.
This extension provides a mechanism for callers to set their system-wide
priority. The default queue priority is
`VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT`.

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

  - [VkDeviceQueueGlobalPriorityCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkQueueGlobalPriorityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_GLOBAL_PRIORITY_EXTENSION_NAME`

- `VK_EXT_GLOBAL_PRIORITY_SPEC_VERSION`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_NOT_PERMITTED_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2017-11-03 (Andres Rodriguez)

  - Fixed VkQueueGlobalPriorityEXT missing \_EXT suffix

- Revision 1, 2017-10-06 (Andres Rodriguez)

  - First version.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_global_priority"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
