# VkPhysicalDeviceInvocationMaskFeaturesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceInvocationMaskFeaturesHUAWEI - Structure describing
invocation mask features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceInvocationMaskFeaturesHUAWEI` structure is defined
as:

``` c
// Provided by VK_HUAWEI_invocation_mask
typedef struct VkPhysicalDeviceInvocationMaskFeaturesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           invocationMask;
} VkPhysicalDeviceInvocationMaskFeaturesHUAWEI;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-invocationMask"></span> `invocationMask` indicates
  that the implementation supports the use of an invocation mask image
  to optimize the ray dispatch.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceInvocationMaskFeaturesHUAWEI` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceInvocationMaskFeaturesHUAWEI` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceInvocationMaskFeaturesHUAWEI-sType-sType"
  id="VUID-VkPhysicalDeviceInvocationMaskFeaturesHUAWEI-sType-sType"></a>
  VUID-VkPhysicalDeviceInvocationMaskFeaturesHUAWEI-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INVOCATION_MASK_FEATURES_HUAWEI`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_invocation_mask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_invocation_mask.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceInvocationMaskFeaturesHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
