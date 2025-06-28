# VkCommandBufferSubmitInfo(3) Manual Page

## Name

VkCommandBufferSubmitInfo - Structure specifying a command buffer submission



## [](#_c_specification)C Specification

The `VkCommandBufferSubmitInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkCommandBufferSubmitInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkCommandBuffer    commandBuffer;
    uint32_t           deviceMask;
} VkCommandBufferSubmitInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_synchronization2
typedef VkCommandBufferSubmitInfo VkCommandBufferSubmitInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `commandBuffer` is a [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) to be submitted for execution.
- `deviceMask` is a bitmask indicating which devices in a device group execute the command buffer. A `deviceMask` of `0` is equivalent to setting all bits corresponding to valid devices in the group to `1`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCommandBufferSubmitInfo-commandBuffer-03890)VUID-VkCommandBufferSubmitInfo-commandBuffer-03890  
  `commandBuffer` **must** not have been allocated with `VK_COMMAND_BUFFER_LEVEL_SECONDARY`
- [](#VUID-VkCommandBufferSubmitInfo-deviceMask-03891)VUID-VkCommandBufferSubmitInfo-deviceMask-03891  
  If `deviceMask` is not `0`, it **must** be a valid device mask
- [](#VUID-VkCommandBufferSubmitInfo-commandBuffer-09445)VUID-VkCommandBufferSubmitInfo-commandBuffer-09445  
  If any render pass instance in `commandBuffer` was recorded with a [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html) structure in its pNext chain and did not specify the `VK_RENDERING_RESUMING_BIT` flag, a [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeSubmitInfoARM.html) **must** be included in the `pNext` chain
- [](#VUID-VkCommandBufferSubmitInfo-pNext-09446)VUID-VkCommandBufferSubmitInfo-pNext-09446  
  If a [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeSubmitInfoARM.html) is included in the `pNext` chain, the value of [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeSubmitInfoARM.html)::`stripeSemaphoreInfoCount` **must** be equal to the sum of the [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html)::`stripeInfoCount` parameters provided to render pass instances recorded in `commandBuffer` that did not specify the `VK_RENDERING_RESUMING_BIT` flag

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferSubmitInfo-sType-sType)VUID-VkCommandBufferSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO`
- [](#VUID-VkCommandBufferSubmitInfo-pNext-pNext)VUID-VkCommandBufferSubmitInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeSubmitInfoARM.html)
- [](#VUID-VkCommandBufferSubmitInfo-sType-unique)VUID-VkCommandBufferSubmitInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkCommandBufferSubmitInfo-commandBuffer-parameter)VUID-VkCommandBufferSubmitInfo-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferSubmitInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0