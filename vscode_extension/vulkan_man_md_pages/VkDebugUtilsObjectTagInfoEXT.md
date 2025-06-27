# VkDebugUtilsObjectTagInfoEXT(3) Manual Page

## Name

VkDebugUtilsObjectTagInfoEXT - Specify parameters of a tag to attach to
an object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDebugUtilsObjectTagInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsObjectTagInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkObjectType       objectType;
    uint64_t           objectHandle;
    uint64_t           tagName;
    size_t             tagSize;
    const void*        pTag;
} VkDebugUtilsObjectTagInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) specifying the
  type of the object to be named.

- `objectHandle` is the object to be tagged.

- `tagName` is a numerical identifier of the tag.

- `tagSize` is the number of bytes of data to attach to the object.

- `pTag` is a pointer to an array of `tagSize` bytes containing the data
  to be associated with the object.

## <a href="#_description" class="anchor"></a>Description

The `tagName` parameter gives a name or identifier to the type of data
being tagged. This can be used by debugging layers to easily filter for
only data that can be used by that implementation.

Valid Usage

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-objectType-01908"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-objectType-01908"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-objectType-01908  
  `objectType` **must** not be `VK_OBJECT_TYPE_UNKNOWN`

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-objectHandle-01910"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-objectHandle-01910"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-objectHandle-01910  
  `objectHandle` **must** be a valid Vulkan handle of the type
  associated with `objectType` as defined in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debugging-object-types"
  target="_blank" rel="noopener"><code>VkObjectType</code> and Vulkan
  Handle Relationship</a> table

Valid Usage (Implicit)

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-sType-sType"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-sType-sType"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT`

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-pNext-pNext"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-pNext-pNext"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-objectType-parameter"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-objectType-parameter"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html)
  value

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-pTag-parameter"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-pTag-parameter"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-pTag-parameter  
  `pTag` **must** be a valid pointer to an array of `tagSize` bytes

- <a href="#VUID-VkDebugUtilsObjectTagInfoEXT-tagSize-arraylength"
  id="VUID-VkDebugUtilsObjectTagInfoEXT-tagSize-arraylength"></a>
  VUID-VkDebugUtilsObjectTagInfoEXT-tagSize-arraylength  
  `tagSize` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSetDebugUtilsObjectTagEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDebugUtilsObjectTagEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsObjectTagInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
