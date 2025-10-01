# VK\_NV\_ray\_tracing\_invocation\_reorder(3) Manual Page

## Name

VK\_NV\_ray\_tracing\_invocation\_reorder - device extension



## [](#_registered_extension_number)Registered Extension Number

491

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_shader\_invocation\_reorder](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_shader_invocation_reorder.html)

## [](#_contact)Contact

- Eric Werness [\[GitHub\]ewerness-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing_invocation_reorder%5D%20%40ewerness-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing_invocation_reorder%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-11-02

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_shader_invocation_reorder`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_shader_invocation_reorder.txt)

**Contributors**

- Eric Werness, NVIDIA
- Ashwin Lele, NVIDIA

## [](#_description)Description

The ray tracing pipeline API provides some ability to reorder for locality, but it is useful to have more control over how the reordering happens and what information is included in the reordering. The shader API provides a hit object to contain result information from the hit which can be used as part of the explicit sorting plus options that contain an integer for hint bits to use to add more locality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkRayTracingInvocationReorderModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingInvocationReorderModeNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_RAY_TRACING_INVOCATION_REORDER_EXTENSION_NAME`
- `VK_NV_RAY_TRACING_INVOCATION_REORDER_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_PROPERTIES_NV`

## [](#_hlsl_mapping)HLSL Mapping

HLSL does not provide this functionality natively yet.

However, it is possible to use this functionality via [SPIR-V Intrinsics](https://github.com/microsoft/DirectXShaderCompiler/wiki/GL_EXT_spirv_intrinsics-for-SPIR-V-code-gen).

The codes for shader invocation reorder are obtained from [this page](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_shader_invocation_reorder.html):

```c
#define ShaderInvocationReorderNV 5383
#define HitObjectAttributeNV 5385

#define OpHitObjectRecordHitMotionNV 5249
#define OpHitObjectRecordHitWithIndexMotionNV 5250
#define OpHitObjectRecordMissMotionNV 5251
#define OpHitObjectGetWorldToObjectNV 5252
#define OpHitObjectGetObjectToWorldNV 5253
#define OpHitObjectGetObjectRayDirectionNV 5254
#define OpHitObjectGetObjectRayOriginNV 5255
#define OpHitObjectTraceRayMotionNV 5256
#define OpHitObjectGetShaderRecordBufferHandleNV 5257
#define OpHitObjectGetShaderBindingTableRecordIndexNV 5258
#define OpHitObjectRecordEmptyNV 5259
#define OpHitObjectTraceRayNV 5260
#define OpHitObjectRecordHitNV 5261
#define OpHitObjectRecordHitWithIndexNV 5262
#define OpHitObjectRecordMissNV 5263
#define OpHitObjectExecuteShaderNV 5264
#define OpHitObjectGetCurrentTimeNV 5265
#define OpHitObjectGetAttributesNV 5266
#define OpHitObjectGetHitKindNV 5267
#define OpHitObjectGetPrimitiveIndexNV 5268
#define OpHitObjectGetGeometryIndexNV 5269
#define OpHitObjectGetInstanceIdNV 5270
#define OpHitObjectGetInstanceCustomIndexNV 5271
#define OpHitObjectGetWorldRayDirectionNV 5272
#define OpHitObjectGetWorldRayOriginNV 5273
#define OpHitObjectGetRayTMaxNV 5274
#define OpHitObjectGetRayTMinNV 5275
#define OpHitObjectIsEmptyNV 5276
#define OpHitObjectIsHitNV 5277
#define OpHitObjectIsMissNV 5278
#define OpReorderThreadWithHitObjectNV 5279
#define OpReorderThreadWithHintNV 5280
#define OpTypeHitObjectNV 5281
```

The capability and extension need to be added:

```c
[[vk::ext_capability(ShaderInvocationReorderNV)]]
[[vk::ext_extension("SPV_NV_shader_invocation_reorder")]]
```

The creation of the `HitObject` type can be done like this:

```c
[[vk::ext_type_def(HitObjectAttributeNV, OpTypeHitObjectNV)]]
void createHitObjectNV();
#define HitObjectNV vk::ext_type<HitObjectAttributeNV>
```

The payload:

- must be global
- needs the `RayPayloadKHR` attribute as an extra storage class

```c
struct [raypayload] HitPayload
{
  float hitT : write(closesthit, miss) : read(caller);
  int instanceIndex : write(closesthit) : read(caller);
  float3 pos : write(closesthit) : read(caller);
  float3 nrm : write(closesthit) : read(caller);
};

#define RayPayloadKHR 5338
[[vk::ext_storage_class(RayPayloadKHR)]] static HitPayload payload;
```

Here is the declaration of a few invocation reordering functions:

```c
[[vk::ext_instruction(OpHitObjectRecordEmptyNV)]]
void hitObjectRecordEmptyNV([[vk::ext_reference]] HitObjectNV hitObject);

[[vk::ext_instruction(OpHitObjectTraceRayNV)]]
void hitObjectTraceRayNV(
    [[vk::ext_reference]] HitObjectNV hitObject,
    RaytracingAccelerationStructure as,
    uint RayFlags,
    uint CullMask,
    uint SBTOffset,
    uint SBTStride,
    uint MissIndex,
    float3 RayOrigin,
    float RayTmin,
    float3 RayDirection,
    float RayTMax,
    [[vk::ext_reference]] [[vk::ext_storage_class(RayPayloadKHR)]] HitPayload payload
  );

[[vk::ext_instruction(OpReorderThreadWithHintNV)]]
void reorderThreadWithHintNV(int Hint, int Bits);

[[vk::ext_instruction(OpReorderThreadWithHitObjectNV)]]
void reorderThreadWithHitObjectNV([[vk::ext_reference]] HitObjectNV hitObject);

[[vk::ext_instruction(OpHitObjectExecuteShaderNV)]]
void hitObjectExecuteShaderNV([[vk::ext_reference]] HitObjectNV hitObject, [[vk::ext_reference]] [[vk::ext_storage_class(RayPayloadKHR)]] HitPayload payload);

[[vk::ext_instruction(OpHitObjectIsHitNV)]]
bool hitObjectIsHitNV([[vk::ext_reference]] HitObjectNV hitObject);
```

Using the function in the code, can be done like this

```c
  if (USE_SER == 1)
  {
    createHitObjectNV();
    HitObjectNV hObj; //  hitObjectNV hObj;
    hitObjectRecordEmptyNV(hObj); //Initialize to an empty hit object
    hitObjectTraceRayNV(hObj, topLevelAS, rayFlags, 0xFF, 0, 0, 0, r.Origin, 0.0, r.Direction, INFINITE, payload);
    reorderThreadWithHitObjectNV(hObj);
    hitObjectExecuteShaderNV(hObj, payload);
  }
```

Note:

- createHitObjectNV() needs to be call at least once. This can be also done in the main entry of the shader.
- Function with a payload parameter, needs to have the payload struct defined before. There are no templated declaration of the function.

## [](#_version_history)Version History

- Revision 1, 2020-09-12 (Eric Werness, Ashwin Lele)
  
  - Initial external release

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_ray_tracing_invocation_reorder).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0