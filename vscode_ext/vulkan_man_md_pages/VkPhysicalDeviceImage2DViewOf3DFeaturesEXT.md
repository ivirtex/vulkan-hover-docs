# VkPhysicalDeviceImage2DViewOf3DFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImage2DViewOf3DFeaturesEXT - Structure describing
whether single-slice 2D views of 3D images can be used in image
descriptors



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImage2DViewOf3DFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_image_2d_view_of_3d
typedef struct VkPhysicalDeviceImage2DViewOf3DFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           image2DViewOf3D;
    VkBool32           sampler2DViewOf3D;
} VkPhysicalDeviceImage2DViewOf3DFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-image2DViewOf3D"></span> `image2DViewOf3D`
  indicates that the implementation supports using a 2D view of a 3D
  image in a descriptor of type `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` if
  the image is created using
  `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT`.

- <span id="features-sampler2DViewOf3D"></span> `sampler2DViewOf3D`
  indicates that the implementation supports using a 2D view of a 3D
  image in a descriptor of type `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` if the image is created
  using `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceImage2DViewOf3DFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImage2DViewOf3DFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceImage2DViewOf3DFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceImage2DViewOf3DFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceImage2DViewOf3DFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_2D_VIEW_OF_3D_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_2d_view_of_3d](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_2d_view_of_3d.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImage2DViewOf3DFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
