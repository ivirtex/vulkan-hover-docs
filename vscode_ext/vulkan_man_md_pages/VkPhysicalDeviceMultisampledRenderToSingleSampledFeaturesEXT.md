# VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT - Structure
describing whether multisampled rendering to single-sampled attachments
is supported



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT`
structure is defined as:

``` c
// Provided by VK_EXT_multisampled_render_to_single_sampled
typedef struct VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           multisampledRenderToSingleSampled;
} VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-multisampledRenderToSingleSampled"></span>
  `multisampledRenderToSingleSampled` indicates that the implementation
  supports multisampled rendering to single-sampled render pass
  attachments.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT` **can**
also be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_multisampled_render_to_single_sampled](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multisampled_render_to_single_sampled.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
