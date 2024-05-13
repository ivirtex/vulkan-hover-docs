# vkMapMemory(3) Manual Page

## Name

vkMapMemory - Map a memory object into application address space



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve a host virtual address pointer to a region of a mappable
memory object, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkMapMemory(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    VkDeviceSize                                offset,
    VkDeviceSize                                size,
    VkMemoryMapFlags                            flags,
    void**                                      ppData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object to be
  mapped.

- `offset` is a zero-based byte offset from the beginning of the memory
  object.

- `size` is the size of the memory range to map, or `VK_WHOLE_SIZE` to
  map from `offset` to the end of the allocation.

- `flags` is a bitmask of
  [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlagBits.html) specifying additional
  parameters of the memory map operation.

- `ppData` is a pointer to a `void*` variable in which a host-accessible
  pointer to the beginning of the mapped range is returned. This pointer
  minus `offset` **must** be aligned to at least
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`minMemoryMapAlignment`.

## <a href="#_description" class="anchor"></a>Description

After a successful call to `vkMapMemory` the memory object `memory` is
considered to be currently *host mapped*.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is an application error to call <code>vkMapMemory</code> on a
memory object that is already <em>host mapped</em>.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>vkMapMemory</code> will fail if the implementation is unable to
allocate an appropriately sized contiguous virtual address range, e.g.
due to virtual address space fragmentation or platform limits. In such
cases, <code>vkMapMemory</code> <strong>must</strong> return
<code>VK_ERROR_MEMORY_MAP_FAILED</code>. The application
<strong>can</strong> improve the likelihood of success by reducing the
size of the mapped range and/or removing unneeded mappings using <a
href="vkUnmapMemory.html">vkUnmapMemory</a>.</p></td>
</tr>
</tbody>
</table>

`vkMapMemory` does not check whether the device memory is currently in
use before returning the host-accessible pointer. The application
**must** guarantee that any previously submitted command that writes to
this range has completed before the host reads from or writes to that
range, and that any previously submitted command that reads from that
range has completed before the host writes to that region (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-host-writes"
target="_blank" rel="noopener">here</a> for details on fulfilling such a
guarantee). If the device memory was allocated without the
`VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` set, these guarantees **must** be
made for an extended range: the application **must** round down the
start of the range to the nearest multiple of
[VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`,
and round the end of the range up to the nearest multiple of
[VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`.

While a range of device memory is host mapped, the application is
responsible for synchronizing both device and host access to that memory
range.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is important for the application developer to become meticulously
familiar with all of the mechanisms described in the chapter on <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
target="_blank" rel="noopener">Synchronization and Cache Control</a> as
they are crucial to maintaining memory access ordering.</p></td>
</tr>
</tbody>
</table>

Calling `vkMapMemory` is equivalent to calling
[vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory2KHR.html) with an empty `pNext` chain.

Valid Usage

- <a href="#VUID-vkMapMemory-memory-00678"
  id="VUID-vkMapMemory-memory-00678"></a>
  VUID-vkMapMemory-memory-00678  
  `memory` **must** not be currently host mapped

- <a href="#VUID-vkMapMemory-offset-00679"
  id="VUID-vkMapMemory-offset-00679"></a>
  VUID-vkMapMemory-offset-00679  
  `offset` **must** be less than the size of `memory`

- <a href="#VUID-vkMapMemory-size-00680"
  id="VUID-vkMapMemory-size-00680"></a> VUID-vkMapMemory-size-00680  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater
  than `0`

- <a href="#VUID-vkMapMemory-size-00681"
  id="VUID-vkMapMemory-size-00681"></a> VUID-vkMapMemory-size-00681  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less
  than or equal to the size of the `memory` minus `offset`

- <a href="#VUID-vkMapMemory-memory-00682"
  id="VUID-vkMapMemory-memory-00682"></a>
  VUID-vkMapMemory-memory-00682  
  `memory` **must** have been created with a memory type that reports
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`

- <a href="#VUID-vkMapMemory-memory-00683"
  id="VUID-vkMapMemory-memory-00683"></a>
  VUID-vkMapMemory-memory-00683  
  `memory` **must** not have been allocated with multiple instances

- <a href="#VUID-vkMapMemory-flags-09568"
  id="VUID-vkMapMemory-flags-09568"></a> VUID-vkMapMemory-flags-09568  
  `VK_MEMORY_MAP_PLACED_BIT_EXT` **must** not be set in `flags`

Valid Usage (Implicit)

- <a href="#VUID-vkMapMemory-device-parameter"
  id="VUID-vkMapMemory-device-parameter"></a>
  VUID-vkMapMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkMapMemory-memory-parameter"
  id="VUID-vkMapMemory-memory-parameter"></a>
  VUID-vkMapMemory-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-vkMapMemory-flags-parameter"
  id="VUID-vkMapMemory-flags-parameter"></a>
  VUID-vkMapMemory-flags-parameter  
  `flags` **must** be a valid combination of
  [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlagBits.html) values

- <a href="#VUID-vkMapMemory-ppData-parameter"
  id="VUID-vkMapMemory-ppData-parameter"></a>
  VUID-vkMapMemory-ppData-parameter  
  `ppData` **must** be a valid pointer to a pointer value

- <a href="#VUID-vkMapMemory-memory-parent"
  id="VUID-vkMapMemory-memory-parent"></a>
  VUID-vkMapMemory-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `memory` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_MEMORY_MAP_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkMapMemory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
