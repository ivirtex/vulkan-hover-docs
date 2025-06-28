# VkDeviceMemory(3) Manual Page

## Name

VkDeviceMemory - Opaque handle to a device memory object



## [](#_c_specification)C Specification

A Vulkan device operates on data in device memory via memory objects that are represented in the API by a `VkDeviceMemory` handle:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDeviceMemory)
```

## [](#_description)Description

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindAccelerationStructureMemoryInfoNV.html), [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html), [VkBindDataGraphPipelineSessionMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDataGraphPipelineSessionMemoryInfoARM.html), [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html), [VkBindTensorMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindTensorMemoryInfoARM.html), [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVideoSessionMemoryInfoKHR.html), [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html), [VkExportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalBufferInfoEXT.html), [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMappedMemoryRange.html), [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html), [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetFdInfoKHR.html), [VkMemoryGetMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetMetalHandleInfoEXT.html), [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetRemoteAddressInfoNV.html), [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetWin32HandleInfoKHR.html), [VkMemoryGetZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetZirconHandleInfoFUCHSIA.html), [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html), [VkMemoryUnmapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfo.html), [VkSparseImageMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBind.html), [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html), [VkTileMemoryBindInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryBindInfoQCOM.html), [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html), [VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html), [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateMemory.html), [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory.html), [vkBindImageMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindImageMemory.html), [vkFreeMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFreeMemory.html), [vkGetDeviceMemoryCommitment](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryCommitment.html), [vkGetMemoryWin32HandleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandleNV.html), [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html), [vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetDeviceMemoryPriorityEXT.html), [vkUnmapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceMemory)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0