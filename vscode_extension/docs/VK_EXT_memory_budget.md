# VK_EXT_memory_budget(3) Manual Page

## Name

VK_EXT_memory_budget - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

238

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_memory_budget%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_memory_budget%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-10-08

**Contributors**  
- Jeff Bolz, NVIDIA

- Jeff Juliano, NVIDIA

## <a href="#_description" class="anchor"></a>Description

While running a Vulkan application, other processes on the machine might
also be attempting to use the same device memory, which can pose
problems. This extension adds support for querying the amount of memory
used and the total memory budget for a memory heap. The values returned
by this query are implementation-dependent and can depend on a variety
of factors including operating system and system load.

The
[VkPhysicalDeviceMemoryBudgetPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html)::`heapBudget`
values can be used as a guideline for how much total memory from each
heap the **current process** can use at any given time, before
allocations may start failing or causing performance degradation. The
values may change based on other activity in the system that is outside
the scope and control of the Vulkan implementation.

The
[VkPhysicalDeviceMemoryBudgetPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html)::`heapUsage`
will display the **current process** estimated heap usage.

With this information, the idea is for an application at some interval
(once per frame, per few seconds, etc) to query `heapBudget` and
`heapUsage`. From here the application can notice if it is over budget
and decide how it wants to handle the memory situation (free it, move to
host memory, changing mipmap levels, etc). This extension is designed to
be used in concert with
[`VK_EXT_memory_priority`](VK_EXT_memory_priority.html) to help with
this part of memory management.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html):

  - [VkPhysicalDeviceMemoryBudgetPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_MEMORY_BUDGET_EXTENSION_NAME`

- `VK_EXT_MEMORY_BUDGET_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-10-08 (Jeff Bolz)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_memory_budget"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
