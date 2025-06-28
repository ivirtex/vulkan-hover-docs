# VkRenderPassStripeBeginInfoARM(3) Manual Page

## Name

VkRenderPassStripeBeginInfoARM - Structure specifying striped rendering information



## [](#_c_specification)C Specification

The `VkRenderPassStripeBeginInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_render_pass_striped
typedef struct VkRenderPassStripeBeginInfoARM {
    VkStructureType                     sType;
    const void*                         pNext;
    uint32_t                            stripeInfoCount;
    const VkRenderPassStripeInfoARM*    pStripeInfos;
} VkRenderPassStripeBeginInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stripeInfoCount` is the number of stripes in this render pass instance
- `pStripeInfos` is a pointer to an array of `stripeInfoCount` [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeInfoARM.html) structures describing the stripes used by the render pass instance.

## [](#_description)Description

This structure can be included in the `pNext` chain of [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) or [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) to define how the render pass instance is split into stripes.

Valid Usage

- [](#VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-09450)VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-09450  
  `stripeInfoCount` **must** be less than or equal to `VkPhysicalDeviceRenderPassStripedPropertiesARM`::`maxRenderPassStripes`
- [](#VUID-VkRenderPassStripeBeginInfoARM-stripeArea-09451)VUID-VkRenderPassStripeBeginInfoARM-stripeArea-09451  
  The `stripeArea` defined by each element of `pStripeInfos` **must** not overlap the `stripeArea` of any other element

Valid Usage (Implicit)

- [](#VUID-VkRenderPassStripeBeginInfoARM-sType-sType)VUID-VkRenderPassStripeBeginInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_BEGIN_INFO_ARM`
- [](#VUID-VkRenderPassStripeBeginInfoARM-pStripeInfos-parameter)VUID-VkRenderPassStripeBeginInfoARM-pStripeInfos-parameter  
  `pStripeInfos` **must** be a valid pointer to an array of `stripeInfoCount` valid [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeInfoARM.html) structures
- [](#VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-arraylength)VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-arraylength  
  `stripeInfoCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_render\_pass\_striped](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_render_pass_striped.html), [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeInfoARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassStripeBeginInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0