# VkCooperativeVectorPropertiesNV(3) Manual Page

## Name

VkCooperativeVectorPropertiesNV - Structure specifying cooperative vector properties



## [](#_c_specification)C Specification

Each `VkCooperativeVectorPropertiesNV` structure describes a single supported combination of types for a matrix-vector multiply (or multiply-add) operation (`OpCooperativeVectorMatrixMulNV` or `OpCooperativeVectorMatrixMulAddNV`).

The `VkCooperativeVectorPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_cooperative_vector
typedef struct VkCooperativeVectorPropertiesNV {
    VkStructureType       sType;
    void*                 pNext;
    VkComponentTypeKHR    inputType;
    VkComponentTypeKHR    inputInterpretation;
    VkComponentTypeKHR    matrixInterpretation;
    VkComponentTypeKHR    biasInterpretation;
    VkComponentTypeKHR    resultType;
    VkBool32              transpose;
} VkCooperativeVectorPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `inputType` is the component type of vector `Input`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `inputInterpretation` is the value of `InputInterpretation`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `matrixInterpretation` is the value of `MatrixInterpretation`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `biasInterpretation` is the value of `BiasInterpretation`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `resultType` is the component type of `Result` `Type`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `transpose` is a boolean indicating whether opaque layout matrices with this combination of input and output types supports transposition.

## [](#_description)Description

`VK_COMPONENT_TYPE_SINT8_PACKED_NV` and `VK_COMPONENT_TYPE_UINT8_PACKED_NV` **must** not be used for members other than `inputInterpretation`.

The following combinations **must** be supported (each row is a required combination):

     inputType inputInterpretation matrixInterpretation biasInterpretation resultType

FLOAT16

FLOAT16

FLOAT16

FLOAT16

FLOAT16

UINT32

SINT8\_PACKED

SINT8

SINT32

SINT32

SINT8

SINT8

SINT8

SINT32

SINT32

FLOAT32

SINT8

SINT8

SINT32

SINT32

FLOAT16

FLOAT\_E4M3

FLOAT\_E4M3

FLOAT16

FLOAT16

FLOAT16

FLOAT\_E5M2

FLOAT\_E5M2

FLOAT16

FLOAT16

Valid Usage (Implicit)

- [](#VUID-VkCooperativeVectorPropertiesNV-sType-sType)VUID-VkCooperativeVectorPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COOPERATIVE_VECTOR_PROPERTIES_NV`
- [](#VUID-VkCooperativeVectorPropertiesNV-pNext-pNext)VUID-VkCooperativeVectorPropertiesNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCooperativeVectorPropertiesNV-inputType-parameter)VUID-VkCooperativeVectorPropertiesNV-inputType-parameter  
  `inputType` **must** be a valid [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html) value
- [](#VUID-VkCooperativeVectorPropertiesNV-inputInterpretation-parameter)VUID-VkCooperativeVectorPropertiesNV-inputInterpretation-parameter  
  `inputInterpretation` **must** be a valid [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html) value
- [](#VUID-VkCooperativeVectorPropertiesNV-matrixInterpretation-parameter)VUID-VkCooperativeVectorPropertiesNV-matrixInterpretation-parameter  
  `matrixInterpretation` **must** be a valid [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html) value
- [](#VUID-VkCooperativeVectorPropertiesNV-biasInterpretation-parameter)VUID-VkCooperativeVectorPropertiesNV-biasInterpretation-parameter  
  `biasInterpretation` **must** be a valid [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html) value
- [](#VUID-VkCooperativeVectorPropertiesNV-resultType-parameter)VUID-VkCooperativeVectorPropertiesNV-resultType-parameter  
  `resultType` **must** be a valid [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html) value

## [](#_see_also)See Also

[VK\_NV\_cooperative\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_vector.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeVectorPropertiesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCooperativeVectorPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0