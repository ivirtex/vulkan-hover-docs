# VkRenderPassStripeInfoARM(3) Manual Page

## Name

VkRenderPassStripeInfoARM - Structure specifying per rendering stripe
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassStripeInfoARM` structure is defined as:

``` c
// Provided by VK_ARM_render_pass_striped
typedef struct VkRenderPassStripeInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkRect2D           stripeArea;
} VkRenderPassStripeInfoARM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stripeArea` is the stripe area, and is described in more detail
  below.

## <a href="#_description" class="anchor"></a>Description

`stripeArea` is the render area that is affected by this stripe of the
render pass instance. It **must** be a subregion of the `renderArea` of
the render pass instance.

Valid Usage

- <a href="#VUID-VkRenderPassStripeInfoARM-stripeArea-09452"
  id="VUID-VkRenderPassStripeInfoARM-stripeArea-09452"></a>
  VUID-VkRenderPassStripeInfoARM-stripeArea-09452  
  `stripeArea.offset.x` **must** be a multiple of
  [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.width`

- <a href="#VUID-VkRenderPassStripeInfoARM-stripeArea-09453"
  id="VUID-VkRenderPassStripeInfoARM-stripeArea-09453"></a>
  VUID-VkRenderPassStripeInfoARM-stripeArea-09453  
  `stripeArea.extent.width` **must** be a multiple of
  [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.width`,
  or the sum of `stripeArea.offset.x` and `stripeArea.extent.width`
  **must** be equal to the `renderArea.extent.width` of the render pass
  instance

- <a href="#VUID-VkRenderPassStripeInfoARM-stripeArea-09454"
  id="VUID-VkRenderPassStripeInfoARM-stripeArea-09454"></a>
  VUID-VkRenderPassStripeInfoARM-stripeArea-09454  
  `stripeArea.offset.y` **must** be a multiple of
  [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.height`

- <a href="#VUID-VkRenderPassStripeInfoARM-stripeArea-09455"
  id="VUID-VkRenderPassStripeInfoARM-stripeArea-09455"></a>
  VUID-VkRenderPassStripeInfoARM-stripeArea-09455  
  `stripeArea.extent.height` **must** be a multiple of
  [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.height`,
  or the sum of `stripeArea.offset.y` and `stripeArea.extent.height`
  **must** be equal to the `renderArea.extent.height` of the render pass
  instance

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassStripeInfoARM-sType-sType"
  id="VUID-VkRenderPassStripeInfoARM-sType-sType"></a>
  VUID-VkRenderPassStripeInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_INFO_ARM`

- <a href="#VUID-VkRenderPassStripeInfoARM-pNext-pNext"
  id="VUID-VkRenderPassStripeInfoARM-pNext-pNext"></a>
  VUID-VkRenderPassStripeInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_render_pass_striped](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_render_pass_striped.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html),
[VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassStripeInfoARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
