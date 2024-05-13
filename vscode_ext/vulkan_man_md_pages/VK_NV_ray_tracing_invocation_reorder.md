# VK_NV_ray_tracing_invocation_reorder(3) Manual Page

## Name

VK_NV_ray_tracing_invocation_reorder - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

491

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_shader_invocation_reorder](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_shader_invocation_reorder.html)

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing_invocation_reorder%5D%20@ewerness-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing_invocation_reorder%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ewerness-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-11-02

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_shader_invocation_reorder`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_shader_invocation_reorder.txt)

**Contributors**  
- Eric Werness, NVIDIA

- Ashwin Lele, NVIDIA

## <a href="#_description" class="anchor"></a>Description

The ray tracing pipeline API provides some ability to reorder for
locality, but it is useful to have more control over how the reordering
happens and what information is included in the reordering. The shader
API provides a hit object to contain result information from the hit
which can be used as part of the explicit sorting plus options that
contain an integer for hint bits to use to add more locality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkRayTracingInvocationReorderModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingInvocationReorderModeNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_RAY_TRACING_INVOCATION_REORDER_EXTENSION_NAME`

- `VK_NV_RAY_TRACING_INVOCATION_REORDER_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_INVOCATION_REORDER_PROPERTIES_NV`

## <a href="#_hlsl_mapping" class="anchor"></a>HLSL Mapping

HLSL does not provide this functionality natively yet.

However, it is possible to use this functionality via [SPIR-V
Intrinsics](https://github.com/microsoft/DirectXShaderCompiler/wiki/GL_EXT_spirv_intrinsics-for-SPIR-V-code-gen).

The codes for shader invocation reorder are obtained from [this
page](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_shader_invocation_reorder.html):

``` c
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

``` c
[[vk::ext_capability(ShaderInvocationReorderNV)]]
[[vk::ext_extension("SPV_NV_shader_invocation_reorder")]]
```

The creation of the `HitObject` type can be done like this:

``` c
[[vk::ext_type_def(HitObjectAttributeNV, OpTypeHitObjectNV)]]
void createHitObjectNV();
#define HitObjectNV vk::ext_type<HitObjectAttributeNV>
```

The payload:

- must be global

- needs the `RayPayloadKHR` attribute as an extra storage class

``` c
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

``` c
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

``` c
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

- createHitObjectNV() needs to be call at least once. This can be also
  done in the main entry of the shader.

- Function with a payload parameter, needs to have the payload struct
  defined before. There are no templated declaration of the function.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-09-12 (Eric Werness, Ashwin Lele)

  - Initial external release

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_ray_tracing_invocation_reorder"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
