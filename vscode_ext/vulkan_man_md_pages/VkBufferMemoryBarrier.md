# VkBufferMemoryBarrier(3) Manual Page

## Name

VkBufferMemoryBarrier - Structure specifying a buffer memory barrier



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferMemoryBarrier` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkBufferMemoryBarrier {
    VkStructureType    sType;
    const void*        pNext;
    VkAccessFlags      srcAccessMask;
    VkAccessFlags      dstAccessMask;
    uint32_t           srcQueueFamilyIndex;
    uint32_t           dstQueueFamilyIndex;
    VkBuffer           buffer;
    VkDeviceSize       offset;
    VkDeviceSize       size;
} VkBufferMemoryBarrier;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">source access mask</a>.

- `dstAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">destination access mask</a>.

- `srcQueueFamilyIndex` is the source queue family for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a>.

- `dstQueueFamilyIndex` is the destination queue family for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a>.

- `buffer` is a handle to the buffer whose backing memory is affected by
  the barrier.

- `offset` is an offset in bytes into the backing memory for `buffer`;
  this is relative to the base offset as bound to the buffer (see
  [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory.html)).

- `size` is a size in bytes of the affected area of backing memory for
  `buffer`, or `VK_WHOLE_SIZE` to use the range from `offset` to the end
  of the buffer.

## <a href="#_description" class="anchor"></a>Description

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to access to
memory through the specified buffer range, via access types in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">source access mask</a> specified by
`srcAccessMask`. If `srcAccessMask` includes `VK_ACCESS_HOST_WRITE_BIT`,
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-available-and-visible"
target="_blank" rel="noopener">memory domain operation</a> is performed
where available memory in the host domain is also made available to the
device domain.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to access to
memory through the specified buffer range, via access types in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">destination access mask</a> specified by
`dstAccessMask`. If `dstAccessMask` includes `VK_ACCESS_HOST_WRITE_BIT`
or `VK_ACCESS_HOST_READ_BIT`, a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-available-and-visible"
target="_blank" rel="noopener">memory domain operation</a> is performed
where available memory in the device domain is also made available to
the host domain.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>When <code>VK_MEMORY_PROPERTY_HOST_COHERENT_BIT</code> is used,
available memory in host domain is automatically made visible to host
domain, and any host write is automatically made available to host
domain.</p></td>
</tr>
</tbody>
</table>

If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, and
`srcQueueFamilyIndex` is equal to the current queue family, then the
memory barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-release"
target="_blank" rel="noopener">queue family release operation</a> for
the specified buffer range, and the second synchronization scope of the
calling command does not apply to this operation.

If `dstQueueFamilyIndex` is not equal to `srcQueueFamilyIndex`, and
`dstQueueFamilyIndex` is equal to the current queue family, then the
memory barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-acquire"
target="_blank" rel="noopener">queue family acquire operation</a> for
the specified buffer range, and the first synchronization scope of the
calling command does not apply to this operation.

Valid Usage

- <a href="#VUID-VkBufferMemoryBarrier-offset-01187"
  id="VUID-VkBufferMemoryBarrier-offset-01187"></a>
  VUID-VkBufferMemoryBarrier-offset-01187  
  `offset` **must** be less than the size of `buffer`

- <a href="#VUID-VkBufferMemoryBarrier-size-01188"
  id="VUID-VkBufferMemoryBarrier-size-01188"></a>
  VUID-VkBufferMemoryBarrier-size-01188  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater
  than `0`

- <a href="#VUID-VkBufferMemoryBarrier-size-01189"
  id="VUID-VkBufferMemoryBarrier-size-01189"></a>
  VUID-VkBufferMemoryBarrier-size-01189  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less
  than or equal to than the size of `buffer` minus `offset`

- <a href="#VUID-VkBufferMemoryBarrier-buffer-01931"
  id="VUID-VkBufferMemoryBarrier-buffer-01931"></a>
  VUID-VkBufferMemoryBarrier-buffer-01931  
  If `buffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkBufferMemoryBarrier-buffer-09095"
  id="VUID-VkBufferMemoryBarrier-buffer-09095"></a>
  VUID-VkBufferMemoryBarrier-buffer-09095  
  If `buffer` was created with a sharing mode of
  `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid
  queue family

- <a href="#VUID-VkBufferMemoryBarrier-buffer-09096"
  id="VUID-VkBufferMemoryBarrier-buffer-09096"></a>
  VUID-VkBufferMemoryBarrier-buffer-09096  
  If `buffer` was created with a sharing mode of
  `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid
  queue family

- <a href="#VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-04087"
  id="VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-04087"></a>
  VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-04087  
  If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, at
  least one of `srcQueueFamilyIndex` or `dstQueueFamilyIndex` **must**
  not be `VK_QUEUE_FAMILY_EXTERNAL` or `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkBufferMemoryBarrier-None-09097"
  id="VUID-VkBufferMemoryBarrier-None-09097"></a>
  VUID-VkBufferMemoryBarrier-None-09097  
  If the [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) extension
  is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is not greater than or equal
  to Version 1.1, `srcQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkBufferMemoryBarrier-None-09098"
  id="VUID-VkBufferMemoryBarrier-None-09098"></a>
  VUID-VkBufferMemoryBarrier-None-09098  
  If the [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) extension
  is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is not greater than or equal
  to Version 1.1, `dstQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-09099"
  id="VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-09099"></a>
  VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-09099  
  If the [VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)
  extension is not enabled `srcQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkBufferMemoryBarrier-dstQueueFamilyIndex-09100"
  id="VUID-VkBufferMemoryBarrier-dstQueueFamilyIndex-09100"></a>
  VUID-VkBufferMemoryBarrier-dstQueueFamilyIndex-09100  
  If the [VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)
  extension is not enabled `dstQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkBufferMemoryBarrier-None-09049"
  id="VUID-VkBufferMemoryBarrier-None-09049"></a>
  VUID-VkBufferMemoryBarrier-None-09049  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature is not enabled, and `buffer` was created with a sharing mode
  of `VK_SHARING_MODE_CONCURRENT`, at least one of `srcQueueFamilyIndex`
  and `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED`

- <a href="#VUID-VkBufferMemoryBarrier-None-09050"
  id="VUID-VkBufferMemoryBarrier-None-09050"></a>
  VUID-VkBufferMemoryBarrier-None-09050  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature is not enabled, and `buffer` was created with a sharing mode
  of `VK_SHARING_MODE_CONCURRENT`, `srcQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkBufferMemoryBarrier-None-09051"
  id="VUID-VkBufferMemoryBarrier-None-09051"></a>
  VUID-VkBufferMemoryBarrier-None-09051  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature is not enabled, and `buffer` was created with a sharing mode
  of `VK_SHARING_MODE_CONCURRENT`, `dstQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`

Valid Usage (Implicit)

- <a href="#VUID-VkBufferMemoryBarrier-sType-sType"
  id="VUID-VkBufferMemoryBarrier-sType-sType"></a>
  VUID-VkBufferMemoryBarrier-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER`

- <a href="#VUID-VkBufferMemoryBarrier-pNext-pNext"
  id="VUID-VkBufferMemoryBarrier-pNext-pNext"></a>
  VUID-VkBufferMemoryBarrier-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)

- <a href="#VUID-VkBufferMemoryBarrier-sType-unique"
  id="VUID-VkBufferMemoryBarrier-sType-unique"></a>
  VUID-VkBufferMemoryBarrier-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBufferMemoryBarrier-buffer-parameter"
  id="VUID-VkBufferMemoryBarrier-buffer-parameter"></a>
  VUID-VkBufferMemoryBarrier-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html),
[vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferMemoryBarrier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
