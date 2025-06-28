# VkPhysicalDeviceShaderIntegerDotProductProperties(3) Manual Page

## Name

VkPhysicalDeviceShaderIntegerDotProductProperties - Structure containing information about integer dot product support for a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderIntegerDotProductProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceShaderIntegerDotProductProperties {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           integerDotProduct8BitUnsignedAccelerated;
    VkBool32           integerDotProduct8BitSignedAccelerated;
    VkBool32           integerDotProduct8BitMixedSignednessAccelerated;
    VkBool32           integerDotProduct4x8BitPackedUnsignedAccelerated;
    VkBool32           integerDotProduct4x8BitPackedSignedAccelerated;
    VkBool32           integerDotProduct4x8BitPackedMixedSignednessAccelerated;
    VkBool32           integerDotProduct16BitUnsignedAccelerated;
    VkBool32           integerDotProduct16BitSignedAccelerated;
    VkBool32           integerDotProduct16BitMixedSignednessAccelerated;
    VkBool32           integerDotProduct32BitUnsignedAccelerated;
    VkBool32           integerDotProduct32BitSignedAccelerated;
    VkBool32           integerDotProduct32BitMixedSignednessAccelerated;
    VkBool32           integerDotProduct64BitUnsignedAccelerated;
    VkBool32           integerDotProduct64BitSignedAccelerated;
    VkBool32           integerDotProduct64BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating8BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating8BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating16BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating16BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating32BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating32BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating64BitUnsignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating64BitSignedAccelerated;
    VkBool32           integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated;
} VkPhysicalDeviceShaderIntegerDotProductProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_shader_integer_dot_product
typedef VkPhysicalDeviceShaderIntegerDotProductProperties VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `integerDotProduct8BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct8BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct8BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct4x8BitPackedUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned dot product operations from operands packed into 32-bit integers using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct4x8BitPackedSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed dot product operations from operands packed into 32-bit integers using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct4x8BitPackedMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness dot product operations from operands packed into 32-bit integers using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct16BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct16BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct16BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct32BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct32BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct32BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct64BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct64BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct64BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating8BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating8BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned accumulating saturating dot product operations from operands packed into 32-bit integers using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed accumulating saturating dot product operations from operands packed into 32-bit integers using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness accumulating saturating dot product operations from operands packed into 32-bit integers using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating16BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating16BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating32BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating32BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating64BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating64BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-integer-dot-product-accelerated).

If the `VkPhysicalDeviceShaderIntegerDotProductProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the integer dot product acceleration information of a physical device.

Note

A dot product operation is deemed accelerated if its implementation provides a performance advantage over application-provided code composed from elementary instructions and/or other dot product instructions, either because the implementation uses optimized machine code sequences whose generation from application-provided code cannot be guaranteed or because it uses hardware features that cannot otherwise be targeted from application-provided code.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderIntegerDotProductProperties-sType-sType)VUID-VkPhysicalDeviceShaderIntegerDotProductProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_PROPERTIES`

## [](#_see_also)See Also

[VK\_KHR\_shader\_integer\_dot\_product](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_integer_dot_product.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderIntegerDotProductProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0