# VkDeviceGroupSubmitInfo(3) Manual Page

## Name

VkDeviceGroupSubmitInfo - Structure indicating which physical devices execute semaphore operations and command buffers



## [](#_c_specification)C Specification

If the `pNext` chain of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html) includes a `VkDeviceGroupSubmitInfo` structure, then that structure includes device indices and masks specifying which physical devices execute semaphore operations and command buffers.

The `VkDeviceGroupSubmitInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupSubmitInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           waitSemaphoreCount;
    const uint32_t*    pWaitSemaphoreDeviceIndices;
    uint32_t           commandBufferCount;
    const uint32_t*    pCommandBufferDeviceMasks;
    uint32_t           signalSemaphoreCount;
    const uint32_t*    pSignalSemaphoreDeviceIndices;
} VkDeviceGroupSubmitInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_device_group
typedef VkDeviceGroupSubmitInfo VkDeviceGroupSubmitInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `waitSemaphoreCount` is the number of elements in the `pWaitSemaphoreDeviceIndices` array.
- `pWaitSemaphoreDeviceIndices` is a pointer to an array of `waitSemaphoreCount` device indices indicating which physical device executes the semaphore wait operation in the corresponding element of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pWaitSemaphores`.
- `commandBufferCount` is the number of elements in the `pCommandBufferDeviceMasks` array.
- `pCommandBufferDeviceMasks` is a pointer to an array of `commandBufferCount` device masks indicating which physical devices execute the command buffer in the corresponding element of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pCommandBuffers`. A physical device executes the command buffer if the corresponding bit is set in the mask.
- `signalSemaphoreCount` is the number of elements in the `pSignalSemaphoreDeviceIndices` array.
- `pSignalSemaphoreDeviceIndices` is a pointer to an array of `signalSemaphoreCount` device indices indicating which physical device executes the semaphore signal operation in the corresponding element of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pSignalSemaphores`.

## [](#_description)Description

If this structure is not present, semaphore operations and command buffers execute on device index zero.

Valid Usage

- [](#VUID-VkDeviceGroupSubmitInfo-waitSemaphoreCount-00082)VUID-VkDeviceGroupSubmitInfo-waitSemaphoreCount-00082  
  `waitSemaphoreCount` **must** equal [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`waitSemaphoreCount`
- [](#VUID-VkDeviceGroupSubmitInfo-commandBufferCount-00083)VUID-VkDeviceGroupSubmitInfo-commandBufferCount-00083  
  `commandBufferCount` **must** equal [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`commandBufferCount`
- [](#VUID-VkDeviceGroupSubmitInfo-signalSemaphoreCount-00084)VUID-VkDeviceGroupSubmitInfo-signalSemaphoreCount-00084  
  `signalSemaphoreCount` **must** equal [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`signalSemaphoreCount`
- [](#VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-00085)VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-00085  
  All elements of `pWaitSemaphoreDeviceIndices` and `pSignalSemaphoreDeviceIndices` **must** be valid device indices
- [](#VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-00086)VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-00086  
  All elements of `pCommandBufferDeviceMasks` **must** be valid device masks

Valid Usage (Implicit)

- [](#VUID-VkDeviceGroupSubmitInfo-sType-sType)VUID-VkDeviceGroupSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO`
- [](#VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-parameter)VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphoreDeviceIndices` **must** be a valid pointer to an array of `waitSemaphoreCount` `uint32_t` values
- [](#VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-parameter)VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-parameter  
  If `commandBufferCount` is not `0`, `pCommandBufferDeviceMasks` **must** be a valid pointer to an array of `commandBufferCount` `uint32_t` values
- [](#VUID-VkDeviceGroupSubmitInfo-pSignalSemaphoreDeviceIndices-parameter)VUID-VkDeviceGroupSubmitInfo-pSignalSemaphoreDeviceIndices-parameter  
  If `signalSemaphoreCount` is not `0`, `pSignalSemaphoreDeviceIndices` **must** be a valid pointer to an array of `signalSemaphoreCount` `uint32_t` values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceGroupSubmitInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0