# VkPipelineExecutableInternalRepresentationKHR(3) Manual Page

## Name

VkPipelineExecutableInternalRepresentationKHR - Structure describing the
textual form of a pipeline executable internal representation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineExecutableInternalRepresentationKHR` structure is defined
as:

``` c
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineExecutableInternalRepresentationKHR {
    VkStructureType    sType;
    void*              pNext;
    char               name[VK_MAX_DESCRIPTION_SIZE];
    char               description[VK_MAX_DESCRIPTION_SIZE];
    VkBool32           isText;
    size_t             dataSize;
    void*              pData;
} VkPipelineExecutableInternalRepresentationKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `name` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a
  null-terminated UTF-8 string which is a short human readable name for
  this internal representation.

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which is a human readable
  description for this internal representation.

- `isText` specifies whether the returned data is text or opaque data.
  If `isText` is `VK_TRUE` then the data returned in `pData` is text and
  is guaranteed to be a null-terminated UTF-8 string.

- `dataSize` is an integer related to the size, in bytes, of the
  internal representation’s data, as described below.

- `pData` is either `NULL` or a pointer to a block of data into which
  the implementation will write the internal representation.

## <a href="#_description" class="anchor"></a>Description

If `pData` is `NULL`, then the size, in bytes, of the internal
representation data is returned in `dataSize`. Otherwise, `dataSize`
must be the size of the buffer, in bytes, pointed to by `pData` and on
return `dataSize` is overwritten with the number of bytes of data
actually written to `pData` including any trailing null character. If
`dataSize` is less than the size, in bytes, of the internal
representation’s data, at most `dataSize` bytes of data will be written
to `pData`, and `VK_INCOMPLETE` will be returned instead of
`VK_SUCCESS`, to indicate that not all the available representation was
returned.

If `isText` is `VK_TRUE` and `pData` is not `NULL` and `dataSize` is not
zero, the last byte written to `pData` will be a null character.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineExecutableInternalRepresentationKHR-sType-sType"
  id="VUID-VkPipelineExecutableInternalRepresentationKHR-sType-sType"></a>
  VUID-VkPipelineExecutableInternalRepresentationKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INTERNAL_REPRESENTATION_KHR`

- <a
  href="#VUID-VkPipelineExecutableInternalRepresentationKHR-pNext-pNext"
  id="VUID-VkPipelineExecutableInternalRepresentationKHR-pNext-pNext"></a>
  VUID-VkPipelineExecutableInternalRepresentationKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPipelineExecutableInternalRepresentationsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableInternalRepresentationsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineExecutableInternalRepresentationKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
