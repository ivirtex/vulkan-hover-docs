# VkRenderPassStripeInfoARM(3) Manual Page

## Name

VkRenderPassStripeInfoARM - Structure specifying per rendering stripe information



## [](#_c_specification)C Specification

The `VkRenderPassStripeInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_render_pass_striped
typedef struct VkRenderPassStripeInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkRect2D           stripeArea;
} VkRenderPassStripeInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stripeArea` is the stripe area, and is described in more detail below.

## [](#_description)Description

`stripeArea` is the render area that is affected by this stripe of the render pass instance. It **must** be a subregion of the `renderArea` of the render pass instance.

Valid Usage

- [](#VUID-VkRenderPassStripeInfoARM-stripeArea-09452)VUID-VkRenderPassStripeInfoARM-stripeArea-09452  
  `stripeArea.offset.x` **must** be a multiple of [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.width`
- [](#VUID-VkRenderPassStripeInfoARM-stripeArea-09453)VUID-VkRenderPassStripeInfoARM-stripeArea-09453  
  `stripeArea.extent.width` **must** be a multiple of [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.width`, or the sum of `stripeArea.offset.x` and `stripeArea.extent.width` **must** be equal to the `renderArea.extent.width` of the render pass instance
- [](#VUID-VkRenderPassStripeInfoARM-stripeArea-09454)VUID-VkRenderPassStripeInfoARM-stripeArea-09454  
  `stripeArea.offset.y` **must** be a multiple of [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.height`
- [](#VUID-VkRenderPassStripeInfoARM-stripeArea-09455)VUID-VkRenderPassStripeInfoARM-stripeArea-09455  
  `stripeArea.extent.height` **must** be a multiple of [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html)::`renderPassStripeGranularity.height`, or the sum of `stripeArea.offset.y` and `stripeArea.extent.height` **must** be equal to the `renderArea.extent.height` of the render pass instance

Valid Usage (Implicit)

- [](#VUID-VkRenderPassStripeInfoARM-sType-sType)VUID-VkRenderPassStripeInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_INFO_ARM`
- [](#VUID-VkRenderPassStripeInfoARM-pNext-pNext)VUID-VkRenderPassStripeInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_ARM\_render\_pass\_striped](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_render_pass_striped.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassStripeInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0