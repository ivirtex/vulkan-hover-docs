# VkRenderPassStripeBeginInfoARM(3) Manual Page

## Name

VkRenderPassStripeBeginInfoARM - Structure specifying striped rendering
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassStripeBeginInfoARM` structure is defined as:

``` c
// Provided by VK_ARM_render_pass_striped
typedef struct VkRenderPassStripeBeginInfoARM {
    VkStructureType                     sType;
    const void*                         pNext;
    uint32_t                            stripeInfoCount;
    const VkRenderPassStripeInfoARM*    pStripeInfos;
} VkRenderPassStripeBeginInfoARM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stripeInfoCount` is the number of stripes in this render pass
  instance

- `pStripeInfos` is a pointer to an array of `stripeInfoCount`
  [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html) structures
  describing the stripes used by the render pass instance.

## <a href="#_description" class="anchor"></a>Description

This structure can be included in the `pNext` chain of
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) or
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) to define how the render pass
instance is split into stripes.

Valid Usage

- <a href="#VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-09450"
  id="VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-09450"></a>
  VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-09450  
  `stripeInfoCount` **must** be less than or equal to
  `VkPhysicalDeviceRenderPassStripedPropertiesARM`::`maxRenderPassStripes`

- <a href="#VUID-VkRenderPassStripeBeginInfoARM-stripeArea-09451"
  id="VUID-VkRenderPassStripeBeginInfoARM-stripeArea-09451"></a>
  VUID-VkRenderPassStripeBeginInfoARM-stripeArea-09451  
  The `stripeArea` defined by each element of `pStripeInfos` **must**
  not overlap the `stripeArea` of any other element

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassStripeBeginInfoARM-sType-sType"
  id="VUID-VkRenderPassStripeBeginInfoARM-sType-sType"></a>
  VUID-VkRenderPassStripeBeginInfoARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_STRIPE_BEGIN_INFO_ARM`

- <a href="#VUID-VkRenderPassStripeBeginInfoARM-pStripeInfos-parameter"
  id="VUID-VkRenderPassStripeBeginInfoARM-pStripeInfos-parameter"></a>
  VUID-VkRenderPassStripeBeginInfoARM-pStripeInfos-parameter  
  `pStripeInfos` **must** be a valid pointer to an array of
  `stripeInfoCount` valid
  [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html) structures

- <a
  href="#VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-arraylength"
  id="VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-arraylength"></a>
  VUID-VkRenderPassStripeBeginInfoARM-stripeInfoCount-arraylength  
  `stripeInfoCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_render_pass_striped](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_render_pass_striped.html),
[VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassStripeBeginInfoARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
