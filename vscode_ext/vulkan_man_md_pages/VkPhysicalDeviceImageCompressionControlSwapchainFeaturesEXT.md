# VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT - Structure
describing whether per-swapchain image compression controls can be
supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT`
structure is defined as:

``` c
// Provided by VK_EXT_image_compression_control_swapchain
typedef struct VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           imageCompressionControlSwapchain;
} VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-imageCompressionControlSwapchain"></span>
  `imageCompressionControlSwapchain` indicates that the implementation
  supports controlling image controls per swapchain and querying image
  compression properties per surface.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT` **can**
also be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_COMPRESSION_CONTROL_SWAPCHAIN_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_compression_control_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control_swapchain.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
