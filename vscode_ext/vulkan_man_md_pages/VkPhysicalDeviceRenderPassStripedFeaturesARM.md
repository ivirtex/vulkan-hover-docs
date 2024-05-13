# VkPhysicalDeviceRenderPassStripedFeaturesARM(3) Manual Page

## Name

VkPhysicalDeviceRenderPassStripedFeaturesARM - Structure describing
whether striped rendering can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRenderPassStripedFeaturesARM` structure is defined
as:

``` c
// Provided by VK_ARM_render_pass_striped
typedef struct VkPhysicalDeviceRenderPassStripedFeaturesARM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           renderPassStriped;
} VkPhysicalDeviceRenderPassStripedFeaturesARM;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceRenderPassStripedFeaturesARM`
structure describe the following features:

## <a href="#_description" class="anchor"></a>Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-renderPassStriped"></span> `renderPassStriped`
  indicates that striped rendering is supported by the implementation.

If the `VkPhysicalDeviceRenderPassStripedFeaturesARM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRenderPassStripedFeaturesARM` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceRenderPassStripedFeaturesARM-sType-sType"
  id="VUID-VkPhysicalDeviceRenderPassStripedFeaturesARM-sType-sType"></a>
  VUID-VkPhysicalDeviceRenderPassStripedFeaturesARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RENDER_PASS_STRIPED_FEATURES_ARM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_render_pass_striped](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_render_pass_striped.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRenderPassStripedFeaturesARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
