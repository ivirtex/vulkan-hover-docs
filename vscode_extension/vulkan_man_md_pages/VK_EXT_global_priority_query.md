# VK_EXT_global_priority_query(3) Manual Page

## Name

VK_EXT_global_priority_query - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

389

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_global_priority.html)  
and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to [VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Yiwei Zhang <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_global_priority_query%5D%20@zhangyiwei%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_global_priority_query%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zhangyiwei</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-03-29

**IP Status**  
No known IP claims.

**Contributors**  
- Yiwei Zhang, Google

## <a href="#_description" class="anchor"></a>Description

This device extension allows applications to query the global queue
priorities supported by a queue family. It allows implementations to
report which global priority levels are treated differently by the
implementation, instead of silently mapping multiple requested global
priority levels to the same internal priority, or using device creation
failure to signal that a requested priority is not supported. It is
intended primarily for use by system integration along with certain
platform-specific priority enforcement rules.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT.html)

- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html):

  - [VkQueueFamilyGlobalPriorityPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyGlobalPriorityPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_GLOBAL_PRIORITY_QUERY_EXTENSION_NAME`

- `VK_EXT_GLOBAL_PRIORITY_QUERY_SPEC_VERSION`

- `VK_MAX_GLOBAL_PRIORITY_SIZE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES_EXT`

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

- Revision 1, 2021-03-29 (Yiwei Zhang)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_global_priority_query"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
