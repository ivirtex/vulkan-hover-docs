# VkPhysicalDeviceImagelessFramebufferFeatures(3) Manual Page

## Name

VkPhysicalDeviceImagelessFramebufferFeatures - Structure indicating
support for imageless framebuffers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImagelessFramebufferFeatures` structure is defined
as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceImagelessFramebufferFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           imagelessFramebuffer;
} VkPhysicalDeviceImagelessFramebufferFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_imageless_framebuffer
typedef VkPhysicalDeviceImagelessFramebufferFeatures VkPhysicalDeviceImagelessFramebufferFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-imagelessFramebuffer"></span>
  `imagelessFramebuffer` indicates that the implementation supports
  specifying the image view for attachments at render pass begin time
  via
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html).

If the `VkPhysicalDeviceImagelessFramebufferFeatures` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImagelessFramebufferFeatures` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImagelessFramebufferFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceImagelessFramebufferFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceImagelessFramebufferFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_imageless_framebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_imageless_framebuffer.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImagelessFramebufferFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
