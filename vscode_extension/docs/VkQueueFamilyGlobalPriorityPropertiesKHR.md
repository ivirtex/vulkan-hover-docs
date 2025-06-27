# VkQueueFamilyGlobalPriorityPropertiesKHR(3) Manual Page

## Name

VkQueueFamilyGlobalPriorityPropertiesKHR - Return structure for queue
family global priority information query



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkQueueFamilyGlobalPriorityPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyGlobalPriorityPropertiesKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_global_priority
typedef struct VkQueueFamilyGlobalPriorityPropertiesKHR {
    VkStructureType             sType;
    void*                       pNext;
    uint32_t                    priorityCount;
    VkQueueGlobalPriorityKHR    priorities[VK_MAX_GLOBAL_PRIORITY_SIZE_KHR];
} VkQueueFamilyGlobalPriorityPropertiesKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_global_priority_query
typedef VkQueueFamilyGlobalPriorityPropertiesKHR VkQueueFamilyGlobalPriorityPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `priorityCount` is the number of supported global queue priorities in
  this queue family, and it **must** be greater than 0.

- `priorities` is an array of `VK_MAX_GLOBAL_PRIORITY_SIZE_KHR`
  [VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html) enums
  representing all supported global queue priorities in this queue
  family. The first `priorityCount` elements of the array will be valid.

## <a href="#_description" class="anchor"></a>Description

If the `VkQueueFamilyGlobalPriorityPropertiesKHR` structure is included
in the `pNext` chain of the
[VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html) structure
passed to
[vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html),
it is filled in with the list of supported global queue priorities for
the indicated family.

The valid elements of `priorities` **must** not contain any duplicate
values.

The valid elements of `priorities` **must** be a continuous sequence of
[VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html) enums in the
ascending order.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For example, returning <code>priorityCount</code> as 3 with supported
<code>priorities</code> as
<code>VK_QUEUE_GLOBAL_PRIORITY_LOW_KHR</code>,
<code>VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR</code> and
<code>VK_QUEUE_GLOBAL_PRIORITY_REALTIME_KHR</code> is not
allowed.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkQueueFamilyGlobalPriorityPropertiesKHR-sType-sType"
  id="VUID-VkQueueFamilyGlobalPriorityPropertiesKHR-sType-sType"></a>
  VUID-VkQueueFamilyGlobalPriorityPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html),
[VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFamilyGlobalPriorityPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
