# VkOpticalFlowSessionBindingPointNV(3) Manual Page

## Name

VkOpticalFlowSessionBindingPointNV - Binding points of an optical flow session



## [](#_c_specification)C Specification

The optical flow session binding points are defined with the following:

```c++
// Provided by VK_NV_optical_flow
typedef enum VkOpticalFlowSessionBindingPointNV {
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_UNKNOWN_NV = 0,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_INPUT_NV = 1,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_REFERENCE_NV = 2,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_HINT_NV = 3,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_FLOW_VECTOR_NV = 4,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_BACKWARD_FLOW_VECTOR_NV = 5,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_COST_NV = 6,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_BACKWARD_COST_NV = 7,
    VK_OPTICAL_FLOW_SESSION_BINDING_POINT_GLOBAL_FLOW_NV = 8,
} VkOpticalFlowSessionBindingPointNV;
```

## [](#_description)Description

- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_INPUT_NV` specifies the binding point for the input frame.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_REFERENCE_NV` specifies the binding point for the input reference frame.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_HINT_NV` specifies the binding point for the optional external hint flow vectors.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_FLOW_VECTOR_NV` specifies the binding point for output flow vectors of default forward flow calculation.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_BACKWARD_FLOW_VECTOR_NV` specifies the binding point for the optional output flow vector map of optional backward flow calculation.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_COST_NV` specifies the binding point for the optional output cost map of default forward flow calculation.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_BACKWARD_COST_NV` specifies the binding point for the optional output cost map of optional backward flow calculation.
- `VK_OPTICAL_FLOW_SESSION_BINDING_POINT_GLOBAL_FLOW_NV` specifies the binding point for the optional global flow value of default forward flow calculation.

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [vkBindOpticalFlowSessionImageNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindOpticalFlowSessionImageNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowSessionBindingPointNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0