# VkRenderPassStripeSubmitInfoARM(3) Manual Page

## Name

VkRenderPassStripeSubmitInfoARM - Structure specifying striped rendering submit information



## [](#_c_specification)C Specification

The `VkRenderPassStripeSubmitInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_render_pass_striped
typedef struct VkRenderPassStripeSubmitInfoARM {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        stripeSemaphoreInfoCount;
    const VkSemaphoreSubmitInfo*    pStripeSemaphoreInfos;
} VkRenderPassStripeSubmitInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stripeSemaphoreInfoCount` is the number of semaphores used to signal stripe completion in the render pass instances in the submitted command buffer.
- `pStripeSemaphoreInfos` is a pointer to an array of `stripeSemaphoreInfoCount` [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html) structures describing the semaphores used to signal stripe completion.

## [](#_description)Description

This structure can be included in the `pNext` chain of [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html) to provide a set of semaphores to be signaled for each striped render pass instance.

The elements of `pStripeSemaphoreInfos` are mapped to render pass instances in [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html)::`commandBuffer` in submission order and in stripe order within each render pass instance. Each semaphore in `pStripeSemaphoreInfos` is signaled when the implementation has completed execution of the associated stripe. In a render pass instance that has multiview enabled, the stripe includes all views in the view mask. In a render pass instance with `layerCount` greater than 1, the stripe includes all layers.

Render pass instances that specify the `VK_RENDERING_RESUMING_BIT` will not have any elements of `pStripeSemaphoreInfos` mapped to them. Instead, for suspending and resuming render pass instances, this mapping is done for the first suspending render pass instance, and the per-stripe semaphores are only signaled for the last resuming render pass instance.

Valid Usage

- [](#VUID-VkRenderPassStripeSubmitInfoARM-semaphore-09447)VUID-VkRenderPassStripeSubmitInfoARM-semaphore-09447  
  The `semaphore` member of each element of `pStripeSemaphoreInfos` **must** have been created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- [](#VUID-VkRenderPassStripeSubmitInfoARM-sType-sType)VUID-VkRenderPassStripeSubmitInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_SUBMIT_INFO_ARM`
- [](#VUID-VkRenderPassStripeSubmitInfoARM-pStripeSemaphoreInfos-parameter)VUID-VkRenderPassStripeSubmitInfoARM-pStripeSemaphoreInfos-parameter  
  `pStripeSemaphoreInfos` **must** be a valid pointer to an array of `stripeSemaphoreInfoCount` valid [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html) structures
- [](#VUID-VkRenderPassStripeSubmitInfoARM-stripeSemaphoreInfoCount-arraylength)VUID-VkRenderPassStripeSubmitInfoARM-stripeSemaphoreInfoCount-arraylength  
  `stripeSemaphoreInfoCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_render\_pass\_striped](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_render_pass_striped.html), [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassStripeSubmitInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0