# VkExternalMemoryAcquireUnmodifiedEXT(3) Manual Page

## Name

VkExternalMemoryAcquireUnmodifiedEXT - Structure specifying that
external memory has remained unmodified since releasing ownership



## <a href="#_c_specification" class="anchor"></a>C Specification

An *acquire operation* **may** have a performance penalty when acquiring
ownership of a subresource range from one of the special queue families
reserved for external memory ownership transfers described above. The
application **can** reduce the performance penalty in some cases by
adding a
[VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)
structure to the `pNext` chain of the *acquire operation*'s memory
barrier structure.

The `VkExternalMemoryAcquireUnmodifiedEXT` structure is defined as:

``` c
// Provided by VK_EXT_external_memory_acquire_unmodified
typedef struct VkExternalMemoryAcquireUnmodifiedEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           acquireUnmodifiedMemory;
} VkExternalMemoryAcquireUnmodifiedEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `acquireUnmodifiedMemory` specifies, if `VK_TRUE`, that no range of
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) bound to the resource of the
  memory barrier’s subresource range was modified at any time since the
  resource’s most recent release of ownership to the queue family
  specified by the memory barrier’s `srcQueueFamilyIndex`. If
  `VK_FALSE`, it specifies nothing.

## <a href="#_description" class="anchor"></a>Description

If the application releases ownership of the subresource range to one of
the special queue families reserved for external memory ownership
transfers with a memory barrier structure, and later re-acquires
ownership from the same queue family with a memory barrier structure,
and if no range of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) bound to the
resource was modified at any time between the *release operation* and
the *acquire operation*, then the application **should** add a
[VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)
structure to the `pNext` chain of the *acquire operation*'s memory
barrier structure because this **may** reduce the performance penalty.

This struct is ignored if `acquireUnmodifiedMemory` is `VK_FALSE`. In
particular, `VK_FALSE` does *not* specify that memory was modified.

This struct is ignored if the memory barrier’s `srcQueueFamilyIndex` is
not a special queue family reserved for external memory ownership
transfers.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The method by which the application determines whether memory was
modified between the <em>release operation</em> and <em>acquire
operation</em> is outside the scope of Vulkan.</p>
<p>For any Vulkan operation that accesses a resource, the application
<strong>must</strong> not assume the implementation accesses the
resource’s memory as read-only, even for <em>apparently</em> read-only
operations such as transfer commands and shader reads.</p>
<p>The validity of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html">VkExternalMemoryAcquireUnmodifiedEXT</a>::<code>acquireUnmodifiedMemory</code>
is independent of memory ranges outside the ranges of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html">VkDeviceMemory</a> bound to the resource. In
particular, it is independent of any implementation-private memory
associated with the resource.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkExternalMemoryAcquireUnmodifiedEXT-acquireUnmodifiedMemory-08922"
  id="VUID-VkExternalMemoryAcquireUnmodifiedEXT-acquireUnmodifiedMemory-08922"></a>
  VUID-VkExternalMemoryAcquireUnmodifiedEXT-acquireUnmodifiedMemory-08922  
  If `acquireUnmodifiedMemory` is `VK_TRUE`, and the memory barrier’s
  `srcQueueFamilyIndex` is a special queue family reserved for external
  memory ownership transfers (as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers</a>),
  then each range of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) bound to the
  resource **must** have remained unmodified during all time since the
  resource’s most recent release of ownership to the queue family

Valid Usage (Implicit)

- <a href="#VUID-VkExternalMemoryAcquireUnmodifiedEXT-sType-sType"
  id="VUID-VkExternalMemoryAcquireUnmodifiedEXT-sType-sType"></a>
  VUID-VkExternalMemoryAcquireUnmodifiedEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_ACQUIRE_UNMODIFIED_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_external_memory_acquire_unmodified](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_external_memory_acquire_unmodified.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalMemoryAcquireUnmodifiedEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
