# VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT -
Structure describing whether rasterization order attachment access can
be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT`
structure is defined as:

``` c
// Provided by VK_EXT_rasterization_order_attachment_access
typedef struct VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rasterizationOrderColorAttachmentAccess;
    VkBool32           rasterizationOrderDepthAttachmentAccess;
    VkBool32           rasterizationOrderStencilAttachmentAccess;
} VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT;
```

or the equivalent

``` c
// Provided by VK_ARM_rasterization_order_attachment_access
typedef VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesARM;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-rasterizationOrderColorAttachmentAccess"></span>
  `rasterizationOrderColorAttachmentAccess` indicates that rasterization
  order access to color and input attachments is supported by the
  implementation.

- <span id="features-rasterizationOrderDepthAttachmentAccess"></span>
  `rasterizationOrderDepthAttachmentAccess` indicates that rasterization
  order access to the depth aspect of depth/stencil and input
  attachments is supported by the implementation.

- <span id="features-rasterizationOrderStencilAttachmentAccess"></span>
  `rasterizationOrderStencilAttachmentAccess` indicates that
  rasterization order access to the stencil aspect of depth/stencil and
  input attachments is supported by the implementation.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT` **can**
also be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_rasterization_order_attachment_access](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_rasterization_order_attachment_access.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
