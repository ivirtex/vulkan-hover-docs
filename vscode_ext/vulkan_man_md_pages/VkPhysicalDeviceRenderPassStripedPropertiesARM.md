# VkPhysicalDeviceRenderPassStripedPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceRenderPassStripedPropertiesARM - Structure describing
striped rendering limits of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRenderPassStripedPropertiesARM` structure is
defined as:

``` c
// Provided by VK_ARM_render_pass_striped
typedef struct VkPhysicalDeviceRenderPassStripedPropertiesARM {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         renderPassStripeGranularity;
    uint32_t           maxRenderPassStripes;
} VkPhysicalDeviceRenderPassStripedPropertiesARM;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceRenderPassStripedPropertiesARM`
structure describe the following limits:

## <a href="#_description" class="anchor"></a>Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-renderPassStripeGranularity"></span>
  `renderPassStripeGranularity` indicates the minimum supported
  granularity of striped render pass regions.

- <span id="limits-maxRenderPassStripes"></span> `maxRenderPassStripes`
  indicates the maximum number of stripes supported in striped
  rendering.

If the `VkPhysicalDeviceRenderPassStripedPropertiesARM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRenderPassStripedPropertiesARM-sType-sType"
  id="VUID-VkPhysicalDeviceRenderPassStripedPropertiesARM-sType-sType"></a>
  VUID-VkPhysicalDeviceRenderPassStripedPropertiesARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_PROPERTIES_ARM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_render_pass_striped](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_render_pass_striped.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRenderPassStripedPropertiesARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
