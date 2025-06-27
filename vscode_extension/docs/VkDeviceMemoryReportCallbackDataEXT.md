# VkDeviceMemoryReportCallbackDataEXT(3) Manual Page

## Name

VkDeviceMemoryReportCallbackDataEXT - Structure specifying parameters
returned to the callback



## <a href="#_c_specification" class="anchor"></a>C Specification

The definition of `VkDeviceMemoryReportCallbackDataEXT` is:

``` c
// Provided by VK_EXT_device_memory_report
typedef struct VkDeviceMemoryReportCallbackDataEXT {
    VkStructureType                     sType;
    void*                               pNext;
    VkDeviceMemoryReportFlagsEXT        flags;
    VkDeviceMemoryReportEventTypeEXT    type;
    uint64_t                            memoryObjectId;
    VkDeviceSize                        size;
    VkObjectType                        objectType;
    uint64_t                            objectHandle;
    uint32_t                            heapIndex;
} VkDeviceMemoryReportCallbackDataEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is 0 and reserved for future use.

- `type` is a
  [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportEventTypeEXT.html)
  type specifying the type of event reported in this
  `VkDeviceMemoryReportCallbackDataEXT` structure.

- `memoryObjectId` is the unique id for the underlying memory object as
  described below.

- `size` is the size of the memory object in bytes. If `type` is
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT`,
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` or
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT`, `size` is
  a valid [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) value. Otherwise, `size` is
  undefined.

- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) value specifying
  the type of the object associated with this device memory report
  event. If `type` is `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT`,
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT`,
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT`,
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT` or
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT`,
  `objectType` is a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) enum.
  Otherwise, `objectType` is undefined.

- `objectHandle` is the object this device memory report event is
  attributed to. If `type` is
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT`,
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT`,
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` or
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT`, `objectHandle` is a
  valid Vulkan handle of the type associated with `objectType` as
  defined in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debugging-object-types"
  target="_blank" rel="noopener"><code>VkObjectType</code> and Vulkan
  Handle Relationship</a> table. Otherwise, `objectHandle` is undefined.

- `heapIndex` describes which memory heap this device memory allocation
  is made from. If `type` is
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT` or
  `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT`,
  `heapIndex` corresponds to one of the valid heaps from the
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)
  structure. Otherwise, `heapIndex` is undefined.

## <a href="#_description" class="anchor"></a>Description

`memoryObjectId` is used to avoid double-counting on the same memory
object.

If an internally-allocated device memory object or a
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) **cannot** be exported,
`memoryObjectId` **must** be unique in the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html).

If an internally-allocated device memory object or a
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) supports being exported,
`memoryObjectId` **must** be unique system wide.

If an internal device memory object or a
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) is backed by an imported external
memory object, `memoryObjectId` **must** be unique system wide.

Implementor’s Note

If the heap backing an internally-allocated device memory **cannot** be
used to back [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), implementations
**can** advertise that heap with no types.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This structure should only be considered valid during the lifetime of
the triggered callback.</p>
<p>For <code>VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT</code> and
<code>VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT</code> events,
<code>objectHandle</code> usually will not yet exist when the
application or tool receives the callback. <code>objectHandle</code>
will only exist when the create or allocate call that triggered the
event returns, and if the allocation or import ends up failing
<code>objectHandle</code> will not ever exist.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceMemoryReportCallbackDataEXT-sType-sType"
  id="VUID-VkDeviceMemoryReportCallbackDataEXT-sType-sType"></a>
  VUID-VkDeviceMemoryReportCallbackDataEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_MEMORY_REPORT_CALLBACK_DATA_EXT`

- <a href="#VUID-VkDeviceMemoryReportCallbackDataEXT-pNext-pNext"
  id="VUID-VkDeviceMemoryReportCallbackDataEXT-pNext-pNext"></a>
  VUID-VkDeviceMemoryReportCallbackDataEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html),
[VK_EXT_device_memory_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_memory_report.html),
[VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportEventTypeEXT.html),
[VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportFlagsEXT.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceMemoryReportCallbackDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
