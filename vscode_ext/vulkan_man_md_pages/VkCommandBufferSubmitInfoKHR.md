# VkCommandBufferSubmitInfo(3) Manual Page

## Name

VkCommandBufferSubmitInfo - Structure specifying a command buffer
submission



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCommandBufferSubmitInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkCommandBufferSubmitInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkCommandBuffer    commandBuffer;
    uint32_t           deviceMask;
} VkCommandBufferSubmitInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_synchronization2
typedef VkCommandBufferSubmitInfo VkCommandBufferSubmitInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `commandBuffer` is a [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) to be
  submitted for execution.

- `deviceMask` is a bitmask indicating which devices in a device group
  execute the command buffer. A `deviceMask` of `0` is equivalent to
  setting all bits corresponding to valid devices in the group to `1`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCommandBufferSubmitInfo-commandBuffer-03890"
  id="VUID-VkCommandBufferSubmitInfo-commandBuffer-03890"></a>
  VUID-VkCommandBufferSubmitInfo-commandBuffer-03890  
  `commandBuffer` **must** not have been allocated with
  `VK_COMMAND_BUFFER_LEVEL_SECONDARY`

- <a href="#VUID-VkCommandBufferSubmitInfo-deviceMask-03891"
  id="VUID-VkCommandBufferSubmitInfo-deviceMask-03891"></a>
  VUID-VkCommandBufferSubmitInfo-deviceMask-03891  
  If `deviceMask` is not `0`, it **must** be a valid device mask

- <a href="#VUID-VkCommandBufferSubmitInfo-commandBuffer-09445"
  id="VUID-VkCommandBufferSubmitInfo-commandBuffer-09445"></a>
  VUID-VkCommandBufferSubmitInfo-commandBuffer-09445  
  If any render pass instance in `commandBuffer` was recorded with a
  [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html)
  structure in its pNext chain and did not specify the
  `VK_RENDERING_RESUMING_BIT` flag, a
  [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeSubmitInfoARM.html)
  **must** be included in the `pNext` chain

- <a href="#VUID-VkCommandBufferSubmitInfo-pNext-09446"
  id="VUID-VkCommandBufferSubmitInfo-pNext-09446"></a>
  VUID-VkCommandBufferSubmitInfo-pNext-09446  
  If a
  [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeSubmitInfoARM.html)
  is included in the `pNext` chain, the value of
  [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeSubmitInfoARM.html)::`stripeSemaphoreInfoCount`
  **must** be equal to the sum of the
  [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html)::`stripeInfoCount`
  parameters provided to render pass instances recorded in
  `commandBuffer` that did not specify the `VK_RENDERING_RESUMING_BIT`
  flag

Valid Usage (Implicit)

- <a href="#VUID-VkCommandBufferSubmitInfo-sType-sType"
  id="VUID-VkCommandBufferSubmitInfo-sType-sType"></a>
  VUID-VkCommandBufferSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO`

- <a href="#VUID-VkCommandBufferSubmitInfo-pNext-pNext"
  id="VUID-VkCommandBufferSubmitInfo-pNext-pNext"></a>
  VUID-VkCommandBufferSubmitInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeSubmitInfoARM.html)

- <a href="#VUID-VkCommandBufferSubmitInfo-sType-unique"
  id="VUID-VkCommandBufferSubmitInfo-sType-unique"></a>
  VUID-VkCommandBufferSubmitInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkCommandBufferSubmitInfo-commandBuffer-parameter"
  id="VUID-VkCommandBufferSubmitInfo-commandBuffer-parameter"></a>
  VUID-VkCommandBufferSubmitInfo-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferSubmitInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
