# VkBindTensorMemoryInfoARM(3) Manual Page

## Name

VkBindTensorMemoryInfoARM - Structure specifying how to bind a tensor to memory



## [](#_c_specification)C Specification

The `VkBindTensorMemoryInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkBindTensorMemoryInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkTensorARM        tensor;
    VkDeviceMemory     memory;
    VkDeviceSize       memoryOffset;
} VkBindTensorMemoryInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensor` is the tensor to be attached to memory.
- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object describing the device memory to attach.
- `memoryOffset` is the start offset of the region of `memory` which is to be bound to the tensor. The number of bytes returned in the `VkMemoryRequirements`::`size` member in `memory`, starting from `memoryOffset` bytes, will be bound to the specified tensor.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindTensorMemoryInfoARM-tensor-09712)VUID-VkBindTensorMemoryInfoARM-tensor-09712  
  `tensor` **must** not already be backed by a memory object
- [](#VUID-VkBindTensorMemoryInfoARM-memoryOffset-09713)VUID-VkBindTensorMemoryInfoARM-memoryOffset-09713  
  `memoryOffset` **must** be less than the size of `memory`
- [](#VUID-VkBindTensorMemoryInfoARM-memory-09714)VUID-VkBindTensorMemoryInfoARM-memory-09714  
  `memory` **must** have been allocated using one of the memory types allowed in the `memoryTypeBits` member of the `VkMemoryRequirements` structure returned from a call to `vkGetTensorMemoryRequirementsARM` with `tensor`
- [](#VUID-VkBindTensorMemoryInfoARM-memoryOffset-09715)VUID-VkBindTensorMemoryInfoARM-memoryOffset-09715  
  `memoryOffset` **must** be an integer multiple of the `alignment` member of the `VkMemoryRequirements` structure returned from a call to `vkGetTensorMemoryRequirementsARM` with `tensor`
- [](#VUID-VkBindTensorMemoryInfoARM-size-09716)VUID-VkBindTensorMemoryInfoARM-size-09716  
  The `size` member of the `VkMemoryRequirements` structure returned from a call to `vkGetTensorMemoryRequirementsARM` with `tensor` **must** be less than or equal to the size of `memory` minus `memoryOffset`
- [](#VUID-VkBindTensorMemoryInfoARM-tensor-09717)VUID-VkBindTensorMemoryInfoARM-tensor-09717  
  If `tensor` requires a dedicated allocation (as reported by [vkGetTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorMemoryRequirementsARM.html) in [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedRequirements.html)::`requiresDedicatedAllocation` for `tensor`), `memory` **must** have been created with [VkMemoryDedicatedAllocateInfoTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfoTensorARM.html)::`tensor` equal to `tensor`
- [](#VUID-VkBindTensorMemoryInfoARM-memory-09806)VUID-VkBindTensorMemoryInfoARM-memory-09806  
  If the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) provided when `memory` was allocated included a [VkMemoryDedicatedAllocateInfoTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfoTensorARM.html) structure in its `pNext` chain, and [VkMemoryDedicatedAllocateInfoTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfoTensorARM.html)::`tensor` was not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), then `tensor` **must** equal [VkMemoryDedicatedAllocateInfoTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfoTensorARM.html)::`tensor`, and `memoryOffset` **must** be zero
- [](#VUID-VkBindTensorMemoryInfoARM-memory-09895)VUID-VkBindTensorMemoryInfoARM-memory-09895  
  If the value of [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` used to allocate `memory` is not `0`, it **must** include at least one of the handles set in [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html)::`handleTypes` when `tensor` was created
- [](#VUID-VkBindTensorMemoryInfoARM-memory-09896)VUID-VkBindTensorMemoryInfoARM-memory-09896  
  If `memory` was allocated by a memory import operation, that is not [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportAndroidHardwareBufferInfoANDROID.html) with a non-`NULL` `buffer` value, the external handle type of the imported memory **must** also have been set in [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html)::`handleTypes` when `tensor` was created
- [](#VUID-VkBindTensorMemoryInfoARM-memory-09897)VUID-VkBindTensorMemoryInfoARM-memory-09897  
  If `memory` was allocated with the [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportAndroidHardwareBufferInfoANDROID.html) memory import operation with a non-`NULL` `buffer` value, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` **must** also have been set in [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html)::`handleTypes` when `tensor` was created
- [](#VUID-VkBindTensorMemoryInfoARM-tensor-09718)VUID-VkBindTensorMemoryInfoARM-tensor-09718  
  If `tensor` was created with the `VK_TENSOR_CREATE_PROTECTED_BIT_ARM` bit set, the tensor **must** be bound to a memory object allocated with a memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`
- [](#VUID-VkBindTensorMemoryInfoARM-tensor-09719)VUID-VkBindTensorMemoryInfoARM-tensor-09719  
  If `tensor` was created with the `VK_TENSOR_CREATE_PROTECTED_BIT_ARM` bit not set, the tensor **must** not be bound to a memory object allocated with a memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

Valid Usage (Implicit)

- [](#VUID-VkBindTensorMemoryInfoARM-sType-sType)VUID-VkBindTensorMemoryInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_TENSOR_MEMORY_INFO_ARM`
- [](#VUID-VkBindTensorMemoryInfoARM-pNext-pNext)VUID-VkBindTensorMemoryInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBindTensorMemoryInfoARM-tensor-parameter)VUID-VkBindTensorMemoryInfoARM-tensor-parameter  
  `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle
- [](#VUID-VkBindTensorMemoryInfoARM-memory-parameter)VUID-VkBindTensorMemoryInfoARM-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkBindTensorMemoryInfoARM-commonparent)VUID-VkBindTensorMemoryInfoARM-commonparent  
  Both of `memory`, and `tensor` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [vkBindTensorMemoryARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindTensorMemoryARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindTensorMemoryInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0