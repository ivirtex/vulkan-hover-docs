# VkBindBufferMemoryInfo(3) Manual Page

## Name

VkBindBufferMemoryInfo - Structure specifying how to bind a buffer to
memory



## <a href="#_c_specification" class="anchor"></a>C Specification

`VkBindBufferMemoryInfo` contains members corresponding to the
parameters of [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory.html).

The `VkBindBufferMemoryInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkBindBufferMemoryInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
    VkDeviceMemory     memory;
    VkDeviceSize       memoryOffset;
} VkBindBufferMemoryInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_bind_memory2
typedef VkBindBufferMemoryInfo VkBindBufferMemoryInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is the buffer to be attached to memory.

- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object describing
  the device memory to attach.

- `memoryOffset` is the start offset of the region of `memory` which is
  to be bound to the buffer. The number of bytes returned in the
  `VkMemoryRequirements`::`size` member in `memory`, starting from
  `memoryOffset` bytes, will be bound to the specified buffer.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-07459"
  id="VUID-VkBindBufferMemoryInfo-buffer-07459"></a>
  VUID-VkBindBufferMemoryInfo-buffer-07459  
  `buffer` **must** not have been bound to a memory object

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-01030"
  id="VUID-VkBindBufferMemoryInfo-buffer-01030"></a>
  VUID-VkBindBufferMemoryInfo-buffer-01030  
  `buffer` **must** not have been created with any sparse memory binding
  flags

- <a href="#VUID-VkBindBufferMemoryInfo-memoryOffset-01031"
  id="VUID-VkBindBufferMemoryInfo-memoryOffset-01031"></a>
  VUID-VkBindBufferMemoryInfo-memoryOffset-01031  
  `memoryOffset` **must** be less than the size of `memory`

- <a href="#VUID-VkBindBufferMemoryInfo-memory-01035"
  id="VUID-VkBindBufferMemoryInfo-memory-01035"></a>
  VUID-VkBindBufferMemoryInfo-memory-01035  
  `memory` **must** have been allocated using one of the memory types
  allowed in the `memoryTypeBits` member of the `VkMemoryRequirements`
  structure returned from a call to `vkGetBufferMemoryRequirements` with
  `buffer`

- <a href="#VUID-VkBindBufferMemoryInfo-memoryOffset-01036"
  id="VUID-VkBindBufferMemoryInfo-memoryOffset-01036"></a>
  VUID-VkBindBufferMemoryInfo-memoryOffset-01036  
  `memoryOffset` **must** be an integer multiple of the `alignment`
  member of the `VkMemoryRequirements` structure returned from a call to
  `vkGetBufferMemoryRequirements` with `buffer`

- <a href="#VUID-VkBindBufferMemoryInfo-size-01037"
  id="VUID-VkBindBufferMemoryInfo-size-01037"></a>
  VUID-VkBindBufferMemoryInfo-size-01037  
  The `size` member of the `VkMemoryRequirements` structure returned
  from a call to `vkGetBufferMemoryRequirements` with `buffer` **must**
  be less than or equal to the size of `memory` minus `memoryOffset`

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-01444"
  id="VUID-VkBindBufferMemoryInfo-buffer-01444"></a>
  VUID-VkBindBufferMemoryInfo-buffer-01444  
  If `buffer` requires a dedicated allocation (as reported by
  [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html)
  in
  [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html)::`requiresDedicatedAllocation`
  for `buffer`), `memory` **must** have been allocated with
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`
  equal to `buffer`

- <a href="#VUID-VkBindBufferMemoryInfo-memory-01508"
  id="VUID-VkBindBufferMemoryInfo-memory-01508"></a>
  VUID-VkBindBufferMemoryInfo-memory-01508  
  If the `VkMemoryAllocateInfo` provided when `memory` was allocated
  included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure in its `pNext` chain, and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `buffer` **must**
  equal
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`,
  and `memoryOffset` **must** be zero

- <a href="#VUID-VkBindBufferMemoryInfo-None-01898"
  id="VUID-VkBindBufferMemoryInfo-None-01898"></a>
  VUID-VkBindBufferMemoryInfo-None-01898  
  If `buffer` was created with the `VK_BUFFER_CREATE_PROTECTED_BIT` bit
  set, the buffer **must** be bound to a memory object allocated with a
  memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-VkBindBufferMemoryInfo-None-01899"
  id="VUID-VkBindBufferMemoryInfo-None-01899"></a>
  VUID-VkBindBufferMemoryInfo-None-01899  
  If `buffer` was created with the `VK_BUFFER_CREATE_PROTECTED_BIT` bit
  not set, the buffer **must** not be bound to a memory object allocated
  with a memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-01038"
  id="VUID-VkBindBufferMemoryInfo-buffer-01038"></a>
  VUID-VkBindBufferMemoryInfo-buffer-01038  
  If `buffer` was created with
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** have been allocated with
  [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html)::`buffer`
  equal to a buffer handle created with identical creation parameters to
  `buffer` and `memoryOffset` **must** be zero

- <a href="#VUID-VkBindBufferMemoryInfo-apiVersion-07920"
  id="VUID-VkBindBufferMemoryInfo-apiVersion-07920"></a>
  VUID-VkBindBufferMemoryInfo-apiVersion-07920  
  If the [VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)
  extension is not enabled,
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and `buffer` was not created with
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** not have been allocated
  dedicated for a specific buffer or image

- <a href="#VUID-VkBindBufferMemoryInfo-memory-02726"
  id="VUID-VkBindBufferMemoryInfo-memory-02726"></a>
  VUID-VkBindBufferMemoryInfo-memory-02726  
  If the value of
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  used to allocate `memory` is not `0`, it **must** include at least one
  of the handles set in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
  when `buffer` was created

- <a href="#VUID-VkBindBufferMemoryInfo-memory-02985"
  id="VUID-VkBindBufferMemoryInfo-memory-02985"></a>
  VUID-VkBindBufferMemoryInfo-memory-02985  
  If `memory` was allocated by a memory import operation, that is not
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  with a non-`NULL` `buffer` value, the external handle type of the
  imported memory **must** also have been set in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
  when `buffer` was created

- <a href="#VUID-VkBindBufferMemoryInfo-memory-02986"
  id="VUID-VkBindBufferMemoryInfo-memory-02986"></a>
  VUID-VkBindBufferMemoryInfo-memory-02986  
  If `memory` was allocated with the
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  memory import operation with a non-`NULL` `buffer` value,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  **must** also have been set in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
  when `buffer` was created

- <a href="#VUID-VkBindBufferMemoryInfo-bufferDeviceAddress-03339"
  id="VUID-VkBindBufferMemoryInfo-bufferDeviceAddress-03339"></a>
  VUID-VkBindBufferMemoryInfo-bufferDeviceAddress-03339  
  If the
  [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)::`bufferDeviceAddress`
  feature is enabled and `buffer` was created with the
  `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` bit set, `memory` **must**
  have been allocated with the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT`
  bit set

- <a
  href="#VUID-VkBindBufferMemoryInfo-bufferDeviceAddressCaptureReplay-09200"
  id="VUID-VkBindBufferMemoryInfo-bufferDeviceAddressCaptureReplay-09200"></a>
  VUID-VkBindBufferMemoryInfo-bufferDeviceAddressCaptureReplay-09200  
  If the
  [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)::`bufferDeviceAddressCaptureReplay`
  feature is enabled and `buffer` was created with the
  `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set, `memory`
  **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-06408"
  id="VUID-VkBindBufferMemoryInfo-buffer-06408"></a>
  VUID-VkBindBufferMemoryInfo-buffer-06408  
  If `buffer` was created with
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
  chained to [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`pNext`,
  `memory` **must** be allocated with a
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)
  chained to [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)::`pNext`

- <a
  href="#VUID-VkBindBufferMemoryInfo-descriptorBufferCaptureReplay-08112"
  id="VUID-VkBindBufferMemoryInfo-descriptorBufferCaptureReplay-08112"></a>
  VUID-VkBindBufferMemoryInfo-descriptorBufferCaptureReplay-08112  
  If the `buffer` was created with the
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` bit set

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-09201"
  id="VUID-VkBindBufferMemoryInfo-buffer-09201"></a>
  VUID-VkBindBufferMemoryInfo-buffer-09201  
  If the `buffer` was created with the
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set

- <a href="#VUID-VkBindBufferMemoryInfo-pNext-01605"
  id="VUID-VkBindBufferMemoryInfo-pNext-01605"></a>
  VUID-VkBindBufferMemoryInfo-pNext-01605  
  If the `pNext` chain includes a
  [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryDeviceGroupInfo.html)
  structure, all instances of `memory` specified by
  [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryDeviceGroupInfo.html)::`pDeviceIndices`
  **must** have been allocated

Valid Usage (Implicit)

- <a href="#VUID-VkBindBufferMemoryInfo-sType-sType"
  id="VUID-VkBindBufferMemoryInfo-sType-sType"></a>
  VUID-VkBindBufferMemoryInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO`

- <a href="#VUID-VkBindBufferMemoryInfo-pNext-pNext"
  id="VUID-VkBindBufferMemoryInfo-pNext-pNext"></a>
  VUID-VkBindBufferMemoryInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryDeviceGroupInfo.html)
  or [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindMemoryStatusKHR.html)

- <a href="#VUID-VkBindBufferMemoryInfo-sType-unique"
  id="VUID-VkBindBufferMemoryInfo-sType-unique"></a>
  VUID-VkBindBufferMemoryInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBindBufferMemoryInfo-buffer-parameter"
  id="VUID-VkBindBufferMemoryInfo-buffer-parameter"></a>
  VUID-VkBindBufferMemoryInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkBindBufferMemoryInfo-memory-parameter"
  id="VUID-VkBindBufferMemoryInfo-memory-parameter"></a>
  VUID-VkBindBufferMemoryInfo-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-VkBindBufferMemoryInfo-commonparent"
  id="VUID-VkBindBufferMemoryInfo-commonparent"></a>
  VUID-VkBindBufferMemoryInfo-commonparent  
  Both of `buffer`, and `memory` **must** have been created, allocated,
  or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html),
[vkBindBufferMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindBufferMemoryInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
