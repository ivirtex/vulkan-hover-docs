# VkPipelineCacheHeaderVersionOne(3) Manual Page

## Name

VkPipelineCacheHeaderVersionOne - Structure describing the layout of the
pipeline cache header



## <a href="#_c_specification" class="anchor"></a>C Specification

Version one of the pipeline cache header is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineCacheHeaderVersionOne {
    uint32_t                        headerSize;
    VkPipelineCacheHeaderVersion    headerVersion;
    uint32_t                        vendorID;
    uint32_t                        deviceID;
    uint8_t                         pipelineCacheUUID[VK_UUID_SIZE];
} VkPipelineCacheHeaderVersionOne;
```

## <a href="#_members" class="anchor"></a>Members

- `headerSize` is the length in bytes of the pipeline cache header.

- `headerVersion` is a
  [VkPipelineCacheHeaderVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersion.html)
  value specifying the version of the header. A consumer of the pipeline
  cache **should** use the cache version to interpret the remainder of
  the cache header.

- `vendorID` is the `VkPhysicalDeviceProperties`::`vendorID` of the
  implementation.

- `deviceID` is the `VkPhysicalDeviceProperties`::`deviceID` of the
  implementation.

- `pipelineCacheUUID` is the
  `VkPhysicalDeviceProperties`::`pipelineCacheUUID` of the
  implementation.

## <a href="#_description" class="anchor"></a>Description

Unlike most structures declared by the Vulkan API, all fields of this
structure are written with the least significant byte first, regardless
of host byte-order.

The C language specification does not define the packing of structure
members. This layout assumes tight structure member packing, with
members laid out in the order listed in the structure, and the intended
size of the structure is 32 bytes. If a compiler produces code that
diverges from that pattern, applications **must** employ another method
to set values at the correct offsets.

Valid Usage

- <a href="#VUID-VkPipelineCacheHeaderVersionOne-headerSize-04967"
  id="VUID-VkPipelineCacheHeaderVersionOne-headerSize-04967"></a>
  VUID-VkPipelineCacheHeaderVersionOne-headerSize-04967  
  `headerSize` **must** be 32

- <a href="#VUID-VkPipelineCacheHeaderVersionOne-headerVersion-04968"
  id="VUID-VkPipelineCacheHeaderVersionOne-headerVersion-04968"></a>
  VUID-VkPipelineCacheHeaderVersionOne-headerVersion-04968  
  `headerVersion` **must** be `VK_PIPELINE_CACHE_HEADER_VERSION_ONE`

- <a href="#VUID-VkPipelineCacheHeaderVersionOne-headerSize-08990"
  id="VUID-VkPipelineCacheHeaderVersionOne-headerSize-08990"></a>
  VUID-VkPipelineCacheHeaderVersionOne-headerSize-08990  
  `headerSize` **must** not exceed the size of the pipeline cache

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCacheHeaderVersionOne-headerVersion-parameter"
  id="VUID-VkPipelineCacheHeaderVersionOne-headerVersion-parameter"></a>
  VUID-VkPipelineCacheHeaderVersionOne-headerVersion-parameter  
  `headerVersion` **must** be a valid
  [VkPipelineCacheHeaderVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersion.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineCacheHeaderVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheHeaderVersion.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCacheHeaderVersionOne"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
