# VkPhysicalDeviceMultiviewFeatures(3) Manual Page

## Name

VkPhysicalDeviceMultiviewFeatures - Structure describing multiview
features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMultiviewFeatures` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceMultiviewFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           multiview;
    VkBool32           multiviewGeometryShader;
    VkBool32           multiviewTessellationShader;
} VkPhysicalDeviceMultiviewFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_multiview
typedef VkPhysicalDeviceMultiviewFeatures VkPhysicalDeviceMultiviewFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-multiview"></span> `multiview` specifies
  whether the implementation supports multiview rendering within a
  render pass. If this feature is not enabled, the view mask of each
  subpass **must** always be zero.

- <span id="extension-features-multiview-gs"></span>
  `multiviewGeometryShader` specifies whether the implementation
  supports multiview rendering within a render pass, with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry"
  target="_blank" rel="noopener">geometry shaders</a>. If this feature
  is not enabled, then a pipeline compiled against a subpass with a
  non-zero view mask **must** not include a geometry shader.

- <span id="extension-features-multiview-tess"></span>
  `multiviewTessellationShader` specifies whether the implementation
  supports multiview rendering within a render pass, with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation"
  target="_blank" rel="noopener">tessellation shaders</a>. If this
  feature is not enabled, then a pipeline compiled against a subpass
  with a non-zero view mask **must** not include any tessellation
  shaders.

If the `VkPhysicalDeviceMultiviewFeatures` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMultiviewFeatures` **can** also be used in the `pNext`
chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively
enable these features.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceMultiviewFeatures-multiviewGeometryShader-00580"
  id="VUID-VkPhysicalDeviceMultiviewFeatures-multiviewGeometryShader-00580"></a>
  VUID-VkPhysicalDeviceMultiviewFeatures-multiviewGeometryShader-00580  
  If `multiviewGeometryShader` is enabled then `multiview` **must** also
  be enabled

- <a
  href="#VUID-VkPhysicalDeviceMultiviewFeatures-multiviewTessellationShader-00581"
  id="VUID-VkPhysicalDeviceMultiviewFeatures-multiviewTessellationShader-00581"></a>
  VUID-VkPhysicalDeviceMultiviewFeatures-multiviewTessellationShader-00581  
  If `multiviewTessellationShader` is enabled then `multiview` **must**
  also be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMultiviewFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceMultiviewFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceMultiviewFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMultiviewFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
