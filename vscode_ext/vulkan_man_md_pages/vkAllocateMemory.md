# vkAllocateMemory(3) Manual Page

## Name

vkAllocateMemory - Allocate device memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To allocate memory objects, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkAllocateMemory(
    VkDevice                                    device,
    const VkMemoryAllocateInfo*                 pAllocateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDeviceMemory*                             pMemory);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `pAllocateInfo` is a pointer to a
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure describing
  parameters of the allocation. A successfully returned allocation
  **must** use the requested parameters — no substitution is permitted
  by the implementation.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pMemory` is a pointer to a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle in which information about the allocated memory is returned.

## <a href="#_description" class="anchor"></a>Description

Allocations returned by `vkAllocateMemory` are guaranteed to meet any
alignment requirement of the implementation. For example, if an
implementation requires 128 byte alignment for images and 64 byte
alignment for buffers, the device memory returned through this mechanism
would be 128-byte aligned. This ensures that applications **can**
correctly suballocate objects of different types (with potentially
different alignment requirements) in the same memory object.

When memory is allocated, its contents are undefined with the following
constraint:

- The contents of unprotected memory **must** not be a function of the
  contents of data protected memory objects, even if those memory
  objects were previously freed.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The contents of memory allocated by one application
<strong>should</strong> not be a function of data from protected memory
objects of another application, even if those memory objects were
previously freed.</p></td>
</tr>
</tbody>
</table>

The maximum number of valid memory allocations that **can** exist
simultaneously within a [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) **may** be restricted
by implementation- or platform-dependent limits. The <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMemoryAllocationCount"
target="_blank" rel="noopener"><code>maxMemoryAllocationCount</code></a>
feature describes the number of allocations that **can** exist
simultaneously before encountering these internal limits.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For historical reasons, if <code>maxMemoryAllocationCount</code> is
exceeded, some implementations may return
<code>VK_ERROR_TOO_MANY_OBJECTS</code>. Exceeding this limit will result
in undefined behavior, and an application should not rely on the use of
the returned error code in order to identify when the limit is
reached.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Many protected memory implementations involve complex hardware and
system software support, and often have additional and much lower limits
on the number of simultaneous protected memory allocations (from memory
types with the <code>VK_MEMORY_PROPERTY_PROTECTED_BIT</code> property)
than for non-protected memory allocations. These limits can be
system-wide, and depend on a variety of factors outside of the Vulkan
implementation, so they cannot be queried in Vulkan. Applications
<strong>should</strong> use as few allocations as possible from such
memory types by suballocating aggressively, and be prepared for
allocation failure even when there is apparently plenty of capacity
remaining in the memory heap. As a guideline, the Vulkan conformance
test suite requires that at least 80 minimum-size allocations can exist
concurrently when no other uses of protected memory are active in the
system.</p></td>
</tr>
</tbody>
</table>

Some platforms **may** have a limit on the maximum size of a single
allocation. For example, certain systems **may** fail to create
allocations with a size greater than or equal to 4GB. Such a limit is
implementation-dependent, and if such a failure occurs then the error
`VK_ERROR_OUT_OF_DEVICE_MEMORY` **must** be returned. This limit is
advertised in
[VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxMemoryAllocationSize`.

The cumulative memory size allocated to a heap **can** be limited by the
size of the specified heap. In such cases, allocated memory is tracked
on a per-device and per-heap basis. Some platforms allow overallocation
into other heaps. The overallocation behavior **can** be specified
through the
[`VK_AMD_memory_overallocation_behavior`](VK_AMD_memory_overallocation_behavior.html)
extension.

If the
[VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html)::`pageableDeviceLocalMemory`
feature is enabled, memory allocations made from a heap that includes
`VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` in
[VkMemoryHeap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeap.html)::`flags` **may** be transparently
moved to host-local memory allowing multiple applications to share
device-local memory. If there is no space left in device-local memory
when this new allocation is made, other allocations **may** be moved out
transparently to make room. The operating system will determine which
allocations to move to device-local memory or host-local memory based on
platform-specific criteria. To help the operating system make good
choices, the application **should** set the appropriate memory priority
with
[VkMemoryPriorityAllocateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPriorityAllocateInfoEXT.html)
and adjust it as necessary with
[vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDeviceMemoryPriorityEXT.html).
Higher priority allocations will moved to device-local memory first.

Memory allocations made on heaps without the
`VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` property will not be transparently
promoted to device-local memory by the operating system.

Valid Usage

- <a href="#VUID-vkAllocateMemory-pAllocateInfo-01713"
  id="VUID-vkAllocateMemory-pAllocateInfo-01713"></a>
  VUID-vkAllocateMemory-pAllocateInfo-01713  
  `pAllocateInfo->allocationSize` **must** be less than or equal to
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryHeaps`\[memindex\].`size`
  where `memindex` =
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryTypes`\[`pAllocateInfo->memoryTypeIndex`\].`heapIndex`
  as returned by
  [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html)
  for the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) that `device` was
  created from

- <a href="#VUID-vkAllocateMemory-pAllocateInfo-01714"
  id="VUID-vkAllocateMemory-pAllocateInfo-01714"></a>
  VUID-vkAllocateMemory-pAllocateInfo-01714  
  `pAllocateInfo->memoryTypeIndex` **must** be less than
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryTypeCount`
  as returned by
  [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html)
  for the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) that `device` was
  created from

- <a href="#VUID-vkAllocateMemory-deviceCoherentMemory-02790"
  id="VUID-vkAllocateMemory-deviceCoherentMemory-02790"></a>
  VUID-vkAllocateMemory-deviceCoherentMemory-02790  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceCoherentMemory"
  target="_blank" rel="noopener"><code>deviceCoherentMemory</code></a>
  feature is not enabled, `pAllocateInfo->memoryTypeIndex` **must** not
  identify a memory type supporting
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`

- <a href="#VUID-vkAllocateMemory-maxMemoryAllocationCount-04101"
  id="VUID-vkAllocateMemory-maxMemoryAllocationCount-04101"></a>
  VUID-vkAllocateMemory-maxMemoryAllocationCount-04101  
  There **must** be less than
  `VkPhysicalDeviceLimits`::`maxMemoryAllocationCount` device memory
  allocations currently allocated on the device

Valid Usage (Implicit)

- <a href="#VUID-vkAllocateMemory-device-parameter"
  id="VUID-vkAllocateMemory-device-parameter"></a>
  VUID-vkAllocateMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkAllocateMemory-pAllocateInfo-parameter"
  id="VUID-vkAllocateMemory-pAllocateInfo-parameter"></a>
  VUID-vkAllocateMemory-pAllocateInfo-parameter  
  `pAllocateInfo` **must** be a valid pointer to a valid
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure

- <a href="#VUID-vkAllocateMemory-pAllocator-parameter"
  id="VUID-vkAllocateMemory-pAllocator-parameter"></a>
  VUID-vkAllocateMemory-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkAllocateMemory-pMemory-parameter"
  id="VUID-vkAllocateMemory-pMemory-parameter"></a>
  VUID-vkAllocateMemory-pMemory-parameter  
  `pMemory` **must** be a valid pointer to a
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAllocateMemory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
