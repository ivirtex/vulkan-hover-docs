# VkDeviceGroupSubmitInfo(3) Manual Page

## Name

VkDeviceGroupSubmitInfo - Structure indicating which physical devices
execute semaphore operations and command buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html) includes a
`VkDeviceGroupSubmitInfo` structure, then that structure includes device
indices and masks specifying which physical devices execute semaphore
operations and command buffers.

The `VkDeviceGroupSubmitInfo` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_device_group
typedef VkDeviceGroupSubmitInfo VkDeviceGroupSubmitInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `waitSemaphoreCount` is the number of elements in the
  `pWaitSemaphoreDeviceIndices` array.

- `pWaitSemaphoreDeviceIndices` is a pointer to an array of
  `waitSemaphoreCount` device indices indicating which physical device
  executes the semaphore wait operation in the corresponding element of
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pWaitSemaphores`.

- `commandBufferCount` is the number of elements in the
  `pCommandBufferDeviceMasks` array.

- `pCommandBufferDeviceMasks` is a pointer to an array of
  `commandBufferCount` device masks indicating which physical devices
  execute the command buffer in the corresponding element of
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pCommandBuffers`. A physical
  device executes the command buffer if the corresponding bit is set in
  the mask.

- `signalSemaphoreCount` is the number of elements in the
  `pSignalSemaphoreDeviceIndices` array.

- `pSignalSemaphoreDeviceIndices` is a pointer to an array of
  `signalSemaphoreCount` device indices indicating which physical device
  executes the semaphore signal operation in the corresponding element
  of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pSignalSemaphores`.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, semaphore operations and command
buffers execute on device index zero.

Valid Usage

- <a href="#VUID-VkDeviceGroupSubmitInfo-waitSemaphoreCount-00082"
  id="VUID-VkDeviceGroupSubmitInfo-waitSemaphoreCount-00082"></a>
  VUID-VkDeviceGroupSubmitInfo-waitSemaphoreCount-00082  
  `waitSemaphoreCount` **must** equal
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`waitSemaphoreCount`

- <a href="#VUID-VkDeviceGroupSubmitInfo-commandBufferCount-00083"
  id="VUID-VkDeviceGroupSubmitInfo-commandBufferCount-00083"></a>
  VUID-VkDeviceGroupSubmitInfo-commandBufferCount-00083  
  `commandBufferCount` **must** equal
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`commandBufferCount`

- <a href="#VUID-VkDeviceGroupSubmitInfo-signalSemaphoreCount-00084"
  id="VUID-VkDeviceGroupSubmitInfo-signalSemaphoreCount-00084"></a>
  VUID-VkDeviceGroupSubmitInfo-signalSemaphoreCount-00084  
  `signalSemaphoreCount` **must** equal
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`signalSemaphoreCount`

- <a
  href="#VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-00085"
  id="VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-00085"></a>
  VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-00085  
  All elements of `pWaitSemaphoreDeviceIndices` and
  `pSignalSemaphoreDeviceIndices` **must** be valid device indices

- <a href="#VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-00086"
  id="VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-00086"></a>
  VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-00086  
  All elements of `pCommandBufferDeviceMasks` **must** be valid device
  masks

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceGroupSubmitInfo-sType-sType"
  id="VUID-VkDeviceGroupSubmitInfo-sType-sType"></a>
  VUID-VkDeviceGroupSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO`

- <a
  href="#VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-parameter"
  id="VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-parameter"></a>
  VUID-VkDeviceGroupSubmitInfo-pWaitSemaphoreDeviceIndices-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphoreDeviceIndices`
  **must** be a valid pointer to an array of `waitSemaphoreCount`
  `uint32_t` values

- <a
  href="#VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-parameter"
  id="VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-parameter"></a>
  VUID-VkDeviceGroupSubmitInfo-pCommandBufferDeviceMasks-parameter  
  If `commandBufferCount` is not `0`, `pCommandBufferDeviceMasks`
  **must** be a valid pointer to an array of `commandBufferCount`
  `uint32_t` values

- <a
  href="#VUID-VkDeviceGroupSubmitInfo-pSignalSemaphoreDeviceIndices-parameter"
  id="VUID-VkDeviceGroupSubmitInfo-pSignalSemaphoreDeviceIndices-parameter"></a>
  VUID-VkDeviceGroupSubmitInfo-pSignalSemaphoreDeviceIndices-parameter  
  If `signalSemaphoreCount` is not `0`, `pSignalSemaphoreDeviceIndices`
  **must** be a valid pointer to an array of `signalSemaphoreCount`
  `uint32_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupSubmitInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
