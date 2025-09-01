# VK\_EXT\_global\_priority\_query(3) Manual Page

## Name

VK\_EXT\_global\_priority\_query - device extension



## [](#_registered_extension_number)Registered Extension Number

389

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_global_priority.html)  
and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Yiwei Zhang [\[GitHub\]zzyiwei](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_global_priority_query%5D%20%40zzyiwei%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_global_priority_query%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-03-29

**IP Status**

No known IP claims.

**Contributors**

- Yiwei Zhang, Google

## [](#_description)Description

This device extension allows applications to query the global queue priorities supported by a queue family. It allows implementations to report which global priority levels are treated differently by the implementation, instead of silently mapping multiple requested global priority levels to the same internal priority, or using device creation failure to signal that a requested priority is not supported. It is intended primarily for use by system integration along with certain platform-specific priority enforcement rules.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyGlobalPriorityPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyGlobalPriorityPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_GLOBAL_PRIORITY_QUERY_EXTENSION_NAME`
- `VK_EXT_GLOBAL_PRIORITY_QUERY_SPEC_VERSION`
- `VK_MAX_GLOBAL_PRIORITY_SIZE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES_EXT`

## [](#_issues)Issues

1\) Can we additionally query whether a caller is permitted to acquire a specific global queue priority in this extension?

**RESOLVED**: No. Whether a caller has enough privilege goes with the OS, and the Vulkan driver cannot really guarantee that the privilege will not change in between this query and the actual queue creation call.

2\) If more than 1 queue using global priority is requested, is there a good way to know which queue is failing the device creation?

**RESOLVED**: No. There is not a good way at this moment, and it is also not quite actionable for the applications to know that because the information may not be accurate. Queue creation can fail because of runtime constraints like insufficient privilege or lack of resource, and the failure is not necessarily tied to that particular queue configuration requested.

## [](#_version_history)Version History

- Revision 1, 2021-03-29 (Yiwei Zhang)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_global_priority_query).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0