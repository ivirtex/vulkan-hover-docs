# VkPipelineCacheHeaderVersionOne(3) Manual Page

## Name

VkPipelineCacheHeaderVersionOne - Structure describing the layout of the pipeline cache header



## [](#_c_specification)C Specification

Version one of the pipeline cache header is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineCacheHeaderVersionOne {
    uint32_t                        headerSize;
    VkPipelineCacheHeaderVersion    headerVersion;
    uint32_t                        vendorID;
    uint32_t                        deviceID;
    uint8_t                         pipelineCacheUUID[VK_UUID_SIZE];
} VkPipelineCacheHeaderVersionOne;
```

## [](#_members)Members

- `headerSize` is the length in bytes of the pipeline cache header.
- `headerVersion` is a [VkPipelineCacheHeaderVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheHeaderVersion.html) value specifying the version of the header. A consumer of the pipeline cache **should** use the cache version to interpret the remainder of the cache header. `headerVersion` **must** be written as exactly 4 bytes.
- `vendorID` is the `VkPhysicalDeviceProperties`::`vendorID` of the implementation.
- `deviceID` is the `VkPhysicalDeviceProperties`::`deviceID` of the implementation.
- `pipelineCacheUUID` is the `VkPhysicalDeviceProperties`::`pipelineCacheUUID` of the implementation.

## [](#_description)Description

Unlike most structures declared by the Vulkan API, all fields of this structure are written with the least significant byte first, regardless of host byte-order.

The C language specification does not define the packing of structure members. This layout assumes tight structure member packing, with members laid out in the order listed in the structure, and the intended size of the structure is 32 bytes. If a compiler produces code that diverges from that pattern, applications **must** employ another method to set values at the correct offsets.

Valid Usage

- [](#VUID-VkPipelineCacheHeaderVersionOne-headerSize-04967)VUID-VkPipelineCacheHeaderVersionOne-headerSize-04967  
  `headerSize` **must** be 32
- [](#VUID-VkPipelineCacheHeaderVersionOne-headerVersion-04968)VUID-VkPipelineCacheHeaderVersionOne-headerVersion-04968  
  `headerVersion` **must** be `VK_PIPELINE_CACHE_HEADER_VERSION_ONE`
- [](#VUID-VkPipelineCacheHeaderVersionOne-headerSize-08990)VUID-VkPipelineCacheHeaderVersionOne-headerSize-08990  
  `headerSize` **must** not exceed the size of the pipeline cache

Valid Usage (Implicit)

- [](#VUID-VkPipelineCacheHeaderVersionOne-headerVersion-parameter)VUID-VkPipelineCacheHeaderVersionOne-headerVersion-parameter  
  `headerVersion` **must** be a valid [VkPipelineCacheHeaderVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheHeaderVersion.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineCacheHeaderVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheHeaderVersion.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCacheHeaderVersionOne)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0