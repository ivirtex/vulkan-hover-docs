# VkAcquireNextImageInfoKHR(3) Manual Page

## Name

VkAcquireNextImageInfoKHR - Structure specifying parameters of the
acquire



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAcquireNextImageInfoKHR` structure is defined as:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkAcquireNextImageInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
    uint64_t           timeout;
    VkSemaphore        semaphore;
    VkFence            fence;
    uint32_t           deviceMask;
} VkAcquireNextImageInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchain` is a non-retired swapchain from which an image is
  acquired.

- `timeout` specifies how long the function waits, in nanoseconds, if no
  image is available.

- `semaphore` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a semaphore to
  signal.

- `fence` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a fence to signal.

- `deviceMask` is a mask of physical devices for which the swapchain
  image will be ready to use when the semaphore or fence is signaled.

## <a href="#_description" class="anchor"></a>Description

If [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html) is used, the
device mask is considered to include all physical devices in the logical
device.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html">vkAcquireNextImage2KHR</a>
signals at most one semaphore, even if the application requests waiting
for multiple physical devices to be ready via the
<code>deviceMask</code>. However, only a single physical device
<strong>can</strong> wait on that semaphore, since the semaphore becomes
unsignaled when the wait succeeds. For other physical devices to wait
for the image to be ready, it is necessary for the application to submit
semaphore signal operation(s) to that first physical device to signal
additional semaphore(s) after the wait succeeds, which the other
physical device(s) <strong>can</strong> wait upon.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkAcquireNextImageInfoKHR-swapchain-01675"
  id="VUID-VkAcquireNextImageInfoKHR-swapchain-01675"></a>
  VUID-VkAcquireNextImageInfoKHR-swapchain-01675  
  `swapchain` **must** not be in the retired state

- <a href="#VUID-VkAcquireNextImageInfoKHR-semaphore-01288"
  id="VUID-VkAcquireNextImageInfoKHR-semaphore-01288"></a>
  VUID-VkAcquireNextImageInfoKHR-semaphore-01288  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it
  **must** be unsignaled

- <a href="#VUID-VkAcquireNextImageInfoKHR-semaphore-01781"
  id="VUID-VkAcquireNextImageInfoKHR-semaphore-01781"></a>
  VUID-VkAcquireNextImageInfoKHR-semaphore-01781  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it
  **must** not have any uncompleted signal or wait operations pending

- <a href="#VUID-VkAcquireNextImageInfoKHR-fence-01289"
  id="VUID-VkAcquireNextImageInfoKHR-fence-01289"></a>
  VUID-VkAcquireNextImageInfoKHR-fence-01289  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it **must** be
  unsignaled and **must** not be associated with any other queue command
  that has not yet completed execution on that queue

- <a href="#VUID-VkAcquireNextImageInfoKHR-semaphore-01782"
  id="VUID-VkAcquireNextImageInfoKHR-semaphore-01782"></a>
  VUID-VkAcquireNextImageInfoKHR-semaphore-01782  
  `semaphore` and `fence` **must** not both be equal to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkAcquireNextImageInfoKHR-deviceMask-01290"
  id="VUID-VkAcquireNextImageInfoKHR-deviceMask-01290"></a>
  VUID-VkAcquireNextImageInfoKHR-deviceMask-01290  
  `deviceMask` **must** be a valid device mask

- <a href="#VUID-VkAcquireNextImageInfoKHR-deviceMask-01291"
  id="VUID-VkAcquireNextImageInfoKHR-deviceMask-01291"></a>
  VUID-VkAcquireNextImageInfoKHR-deviceMask-01291  
  `deviceMask` **must** not be zero

- <a href="#VUID-VkAcquireNextImageInfoKHR-semaphore-03266"
  id="VUID-VkAcquireNextImageInfoKHR-semaphore-03266"></a>
  VUID-VkAcquireNextImageInfoKHR-semaphore-03266  
  `semaphore` **must** have a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- <a href="#VUID-VkAcquireNextImageInfoKHR-sType-sType"
  id="VUID-VkAcquireNextImageInfoKHR-sType-sType"></a>
  VUID-VkAcquireNextImageInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR`

- <a href="#VUID-VkAcquireNextImageInfoKHR-pNext-pNext"
  id="VUID-VkAcquireNextImageInfoKHR-pNext-pNext"></a>
  VUID-VkAcquireNextImageInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkAcquireNextImageInfoKHR-swapchain-parameter"
  id="VUID-VkAcquireNextImageInfoKHR-swapchain-parameter"></a>
  VUID-VkAcquireNextImageInfoKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-VkAcquireNextImageInfoKHR-semaphore-parameter"
  id="VUID-VkAcquireNextImageInfoKHR-semaphore-parameter"></a>
  VUID-VkAcquireNextImageInfoKHR-semaphore-parameter  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkAcquireNextImageInfoKHR-fence-parameter"
  id="VUID-VkAcquireNextImageInfoKHR-fence-parameter"></a>
  VUID-VkAcquireNextImageInfoKHR-fence-parameter  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-VkAcquireNextImageInfoKHR-commonparent"
  id="VUID-VkAcquireNextImageInfoKHR-commonparent"></a>
  VUID-VkAcquireNextImageInfoKHR-commonparent  
  Each of `fence`, `semaphore`, and `swapchain` that are valid handles
  of non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

- Host access to `semaphore` **must** be externally synchronized

- Host access to `fence` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html),
[vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAcquireNextImageInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
