# VkPhysicalDeviceSubpassShadingPropertiesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceSubpassShadingPropertiesHUAWEI - Structure describing
subpass shading properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSubpassShadingPropertiesHUAWEI` structure is
defined as:

``` c
// Provided by VK_HUAWEI_subpass_shading
typedef struct VkPhysicalDeviceSubpassShadingPropertiesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxSubpassShadingWorkgroupSizeAspectRatio;
} VkPhysicalDeviceSubpassShadingPropertiesHUAWEI;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxSubpassShadingWorkgroupSizeAspectRatio"></span>
  `maxSubpassShadingWorkgroupSizeAspectRatio` indicates the maximum
  ratio between the width and height of the portion of the subpass
  shading shader workgroup size.
  `maxSubpassShadingWorkgroupSizeAspectRatio` **must** be a power-of-two
  value, and **must** be less than or equal to max(`WorkgroupSize.x` /
  `WorkgroupSize.y`, `WorkgroupSize.y` / `WorkgroupSize.x`).

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceSubpassShadingPropertiesHUAWEI` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceSubpassShadingPropertiesHUAWEI-sType-sType"
  id="VUID-VkPhysicalDeviceSubpassShadingPropertiesHUAWEI-sType-sType"></a>
  VUID-VkPhysicalDeviceSubpassShadingPropertiesHUAWEI-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_SHADING_PROPERTIES_HUAWEI`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_subpass_shading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_subpass_shading.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSubpassShadingPropertiesHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
