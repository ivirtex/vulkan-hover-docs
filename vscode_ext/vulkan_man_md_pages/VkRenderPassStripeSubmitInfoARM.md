# VkRenderPassStripeSubmitInfoARM(3) Manual Page

## Name

VkRenderPassStripeSubmitInfoARM - Structure specifying striped rendering
submit information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassStripeSubmitInfoARM` structure is defined as:

``` c
// Provided by VK_ARM_render_pass_striped
typedef struct VkRenderPassStripeSubmitInfoARM {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        stripeSemaphoreInfoCount;
    const VkSemaphoreSubmitInfo*    pStripeSemaphoreInfos;
} VkRenderPassStripeSubmitInfoARM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stripeSemaphoreInfoCount` is the number of semaphores used to signal
  stripe completion in the render pass instances in the submitted
  command buffer.

- `pStripeSemaphoreInfos` is a pointer to an array of
  `stripeSemaphoreInfoCount`
  [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html) structures
  describing the semaphores used to signal stripe completion.

## <a href="#_description" class="anchor"></a>Description

This structure can be included in the `pNext` chain of
[VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfo.html) to provide a
set of semaphores to be signaled for each striped render pass instance.

The elements of `pStripeSemaphoreInfos` are mapped to render pass
instances in
[VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfo.html)::`commandBuffer`
in submission order and in stripe order within each render pass
instance. Each semaphore in `pStripeSemaphoreInfos` is signaled when the
implementation has completed execution of the associated stripe. In a
render pass instance that has multiview enabled, the stripe includes all
views in the view mask. In a render pass instance with `layerCount`
greater than 1, the stripe includes all layers.

Valid Usage

- <a href="#VUID-VkRenderPassStripeSubmitInfoARM-semaphore-09447"
  id="VUID-VkRenderPassStripeSubmitInfoARM-semaphore-09447"></a>
  VUID-VkRenderPassStripeSubmitInfoARM-semaphore-09447  
  The `semaphore` member of each element of `pStripeSemaphoreInfos`
  **must** have been created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassStripeSubmitInfoARM-sType-sType"
  id="VUID-VkRenderPassStripeSubmitInfoARM-sType-sType"></a>
  VUID-VkRenderPassStripeSubmitInfoARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_SUBMIT_INFO_ARM`

- <a
  href="#VUID-VkRenderPassStripeSubmitInfoARM-pStripeSemaphoreInfos-parameter"
  id="VUID-VkRenderPassStripeSubmitInfoARM-pStripeSemaphoreInfos-parameter"></a>
  VUID-VkRenderPassStripeSubmitInfoARM-pStripeSemaphoreInfos-parameter  
  `pStripeSemaphoreInfos` **must** be a valid pointer to an array of
  `stripeSemaphoreInfoCount` valid
  [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html) structures

- <a
  href="#VUID-VkRenderPassStripeSubmitInfoARM-stripeSemaphoreInfoCount-arraylength"
  id="VUID-VkRenderPassStripeSubmitInfoARM-stripeSemaphoreInfoCount-arraylength"></a>
  VUID-VkRenderPassStripeSubmitInfoARM-stripeSemaphoreInfoCount-arraylength  
  `stripeSemaphoreInfoCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_render_pass_striped](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_render_pass_striped.html),
[VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassStripeSubmitInfoARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
